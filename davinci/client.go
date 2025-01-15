// Davinci Admin API GO Client
//
// This package is go client to be used for interacting with PingOne DaVinci Administrative APIs.
// Use cases include:
// - Creating Connections
// - Importing Flows
package davinci

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"
)

// baseURL
var baseURL = url.URL{
	Scheme: "https",
	Host:   "orchestrate-api.pingone.com",
	Path:   "/v1",
}

var dvApiHost = map[string]string{
	"NorthAmerica":          "orchestrate-api.pingone.com",
	"Europe":                "orchestrate-api.pingone.eu",
	"AsiaPacific":           "orchestrate-api.pingone.asia",
	"Australia-AsiaPacific": "orchestrate-api.pingone.com.au",
	"Canada":                "orchestrate-api.pingone.ca",
}

var defaultUserAgent = "PingOne-DaVinci-GOLANG-SDK"

var requestMutex sync.Mutex

// const HostURL string = "https://api.singularkey.com/v1"

func (args Params) QueryParams() url.Values {
	q := make(url.Values)

	if args.Page != "" {
		q.Add("page", args.Page)
	}

	if args.Limit != "" {
		q.Add("limit", args.Limit)
	}
	for i, v := range args.ExtraParams {
		q.Add(i, v)
	}

	return q
}

func NewClient(inputs *ClientInput) (*APIClient, error) {

	// adjust host according to received region
	if inputs.PingOneRegion == "" {
		return nil, fmt.Errorf("PingOneRegion must be set")
	} else {
		if dvApiHost[inputs.PingOneRegion] == "" {
			return nil, fmt.Errorf("Invalid region: %v", inputs.PingOneRegion)
		}
		baseURL.Host = dvApiHost[inputs.PingOneRegion]
	}

	hostUrl := baseURL.ResolveReference(&url.URL{}).String()

	if inputs.HostURL != "" {
		hostUrl = inputs.HostURL
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, fmt.Errorf("Got error while creating cookie jar %s", err.Error())
	}
	c := APIClient{
		HTTPClient: &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
			Timeout: 300 * time.Second,
			Jar:     jar},
		HostURL: hostUrl,
	}

	if inputs.AccessToken != "" {
		c.Token = inputs.AccessToken
		c.companyID = inputs.PingOneSSOEnvId
		return &c, nil
	}

	if inputs.Username == "" || inputs.Password == "" {
		// return nil, fmt.Errorf("User or Password not found")
		return &c, nil
	}

	c.Auth = AuthStruct{
		Username: inputs.Username,
		Password: inputs.Password,
	}

	c.UserAgent = defaultUserAgent

	if inputs.UserAgent != "" {
		c.UserAgent = inputs.UserAgent
	}

	// Use P1SSO if available
	if inputs.PingOneSSOEnvId != "" {
		c.PingOneSSOEnvId = inputs.PingOneSSOEnvId
	}
	err = c.DoSignIn(nil)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return &c, nil
}

func (c *APIClient) DoSignIn(targetCompanyId *string) error {
	if c.PingOneSSOEnvId != "" {
		ar, _, err := c.SignInSSOWithResponse(targetCompanyId)
		if err != nil {
			return err
		}

		c.Token = ar.AccessToken
		return nil
	}

	return fmt.Errorf("Sign in failed. Not using SSO")
}

func (c *APIClient) doRequestVerbose(req *http.Request, authToken *string, args *Params) (*DvHttpResponse, *http.Response, error) {
	if authToken != nil {
		token := *authToken
		var bearer = "Bearer " + token
		req.Header.Add("Authorization", bearer)
	}
	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	}
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	if args != nil {
		req.URL.RawQuery = args.QueryParams().Encode()
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, res, err
	}
	defer res.Body.Close()

	rbody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, res, err
	}

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusFound {
		return nil, res, fmt.Errorf("status: %d, body: %s", res.StatusCode, rbody)
	}
	resp := DvHttpResponse{
		Body:       rbody,
		Headers:    res.Header,
		StatusCode: res.StatusCode,
	}

	if res.StatusCode == http.StatusFound && res.Header["Location"] != nil {
		resp.Location, _ = url.Parse(res.Header["Location"][0])
		resp.LocationParams, _ = url.ParseQuery(resp.Location.RawQuery)
		// Handle wepbage hash value strangeness
		if resp.Location.Fragment != "" {
			_, v, ok := strings.Cut(resp.Location.Fragment, "?")
			if ok {
				resp.LocationParams, _ = url.ParseQuery(v)
			}
		}
	}
	if res.Header["Set-Cookie"] != nil {
		c.HTTPClient.Jar.SetCookies(req.URL, res.Cookies())
	}

	return &resp, res, err
}

func (c *APIClient) doRequest(reqIn DvHttpRequest, args *Params) ([]byte, *http.Response, error) {

	slog.Debug("Start doRequest", "Company ID", c.companyID)

	req, err := http.NewRequest(reqIn.Method, reqIn.Url, strings.NewReader(reqIn.Body))
	if err != nil {
		return nil, nil, err
	}

	if c.Token != "" {
		req.Header.Del("Authorization")

		var bearer = "Bearer " + c.Token
		req.Header.Add("Authorization", bearer)
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	}
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	if args != nil {
		req.URL.RawQuery = args.QueryParams().Encode()
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	res.Body = io.NopCloser(bytes.NewBuffer(body))
	if err != nil {
		return nil, res, err
	}

	if res.StatusCode >= 300 {
		var errObject ErrorResponse

		if err := json.Unmarshal(body, &errObject); err != nil {
			return nil, res, err
		}

		b, err := json.Marshal(errObject)
		if err != nil {
			return nil, res, err
		}

		if string(b) != "{}" {
			slog.Error("Error response handled", "response", string(b))
			return nil, res, errObject
		} else {
			slog.Debug("Error response unhandled", "status code", res.StatusCode)
		}
	}

	return body, res, err
}

func (c *APIClient) doRequestRetryable(companyId *string, req DvHttpRequest, args *Params) ([]byte, *http.Response, error) {

	// This API action isn't thread safe - the environment may be switched by another thread.  We need to lock it
	requestMutex.Lock()
	defer requestMutex.Unlock()

	body, res, err := c.exponentialBackOffRetry(func() (any, *http.Response, error) {
		slog.Debug("Start exponentialBackOffRetry", "Company ID", c.companyID)

		// handle environment switching
		if companyId != nil && *companyId != c.companyID {
			_, res, err := c.SetEnvironmentWithResponse(*companyId)
			if err != nil {
				return nil, res, err
			}

			if c.companyID != *companyId {
				return nil, nil, fmt.Errorf("Failed to set environment to %s after successful switch", *companyId)
			}
		}

		return c.doRequest(req, args)
	}, false)
	if err != nil {
		return nil, res, err
	}

	return body.([]byte), res, nil
}

type SDKInterfaceFunc func() (any, *http.Response, error)

var (
	maxRetries               = 10
	maximumRetryAfterBackoff = 30
)

func (c *APIClient) exponentialBackOffRetry(f SDKInterfaceFunc, isAuthCall bool) (interface{}, *http.Response, error) {
	var obj interface{}
	var resp *http.Response
	var err error
	backOffTime := time.Second
	isRetryable, reauthNeeded := false, false

	for i := 0; i < maxRetries; i++ {
		obj, resp, err = f()

		backOffTime, isRetryable, reauthNeeded = testForRetryable(resp, err, backOffTime)

		if isRetryable {
			slog.Warn("Retryable request failed", "retry index", i+1, "error", err, "back off time", backOffTime.String(), "reauth needed", reauthNeeded, "reauth possible", !isAuthCall)
			time.Sleep(backOffTime)

			if reauthNeeded && !isAuthCall {
				slog.Debug("Attempting re-auth")
				err := c.DoSignIn(&c.companyID)
				if err != nil {
					slog.Error("Retry sign in failed", "error", err)
					return obj, resp, err
				}
			}

			continue
		}

		return obj, resp, err
	}

	slog.Warn("Retryable request failed", "Max retries", maxRetries, "error", err)

	return obj, resp, err // output the final error
}

func testForRetryable(r *http.Response, err error, currentBackoff time.Duration) (backoffDuration time.Duration, isRetryable bool, reauthNeeded bool) {

	backoffDuration = currentBackoff

	if r != nil {
		if r.StatusCode == http.StatusNotImplemented || r.StatusCode == http.StatusServiceUnavailable || r.StatusCode == http.StatusTooManyRequests {
			retryAfter, err := parseRetryAfterHeader(r)
			if err != nil {
				slog.Warn("Cannot parse the expected \"Retry-After\" header", "error", err)
				backoffDuration = currentBackoff * 2
			}

			if retryAfter <= time.Duration(maximumRetryAfterBackoff) {
				backoffDuration += time.Duration(maximumRetryAfterBackoff)
			} else {
				backoffDuration += retryAfter
			}
		} else {
			backoffDuration = currentBackoff
		}

		retryAbleCodes := []int{
			http.StatusTooManyRequests,
			http.StatusInternalServerError,
			http.StatusBadGateway,
			http.StatusServiceUnavailable,
			http.StatusGatewayTimeout,
		}

		if slices.Contains(retryAbleCodes, r.StatusCode) {
			slog.Debug("Retryable HTTP status code found", "http status code", r.StatusCode)
			return backoffDuration, true, false
		}
	}

	if err != nil {

		switch t := err.(type) {

		case ErrorResponse:
			if t.HttpResponseCode == http.StatusUnauthorized && t.Code == DV_ERROR_CODE_INVALID_TOKEN_FOR_ENVIRONMENT {
				slog.Warn("Client unauthorized for the environment, available for retry (re-auth needed)", "error", err)
				backoffDuration += (2 * time.Second)
				return backoffDuration, true, true
			}

		default:
			if res1, matchErr := regexp.MatchString(`^http: ContentLength=[0-9]+ with Body length [0-9]+$`, err.Error()); matchErr == nil && res1 {
				slog.Warn("HTTP content error detected, available for retry (re-auth needed)", "error", err)
				backoffDuration += (2 * time.Second)
				return backoffDuration, true, true
			}

			if res1, matchErr := regexp.MatchString(`error\=AuthenticationFailed\&error_description\=unknownError2`, err.Error()); matchErr == nil && res1 {
				slog.Warn("Authentication unknown2 error detected, available for retry", "error", err)
				backoffDuration += (2 * time.Second)
				return backoffDuration, true, false
			}
		}
	}

	return backoffDuration, false, false
}

func parseRetryAfterHeader(resp *http.Response) (time.Duration, error) {
	retryAfterHeader := resp.Header.Get("Retry-After")

	if retryAfterHeader == "" {
		return 0, fmt.Errorf("Retry-After header not found")
	}

	retryAfterSeconds, err := strconv.Atoi(retryAfterHeader)

	if err == nil {
		return time.Duration(retryAfterSeconds) * time.Second, nil
	}

	retryAfterTime, err := http.ParseTime(retryAfterHeader)

	if err != nil {
		return 0, fmt.Errorf("Unable to parse Retry-After header value: %v", err)
	}

	return time.Until(retryAfterTime), nil
}
