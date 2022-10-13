// Davinci Admin API GO Client
//
// This package is go client to be used for interacting with PingOne DaVinci Administrative APIs.
// Use cases include:
// - Creating Connections
// - Importing Flows
package davinci

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
)

// baseURL
var baseURL = url.URL{
	Scheme: "https",
	Host:   "api.singularkey.com",
	Path:   "/v1",
}

// const HostURL string = "https://api.singularkey.com/v1"

type ClientInput struct {
	HostURL   string
	Username  string
	Password  string
	AuthP1SSO AuthP1SSO
}

type Client struct {
	HostURL     string
	HTTPClient  *http.Client
	Token       string
	Auth        AuthStruct
	CompanyID   string
	AuthP1SSO   AuthP1SSO
	AuthRefresh bool
}

type Params struct {
	Page        string
	Limit       string
	ExtraParams map[string]string
	// TODO: figure out what query is
	// query  string
}

type DvHttpRequest struct {
	Method string
	Url    string
	Body   io.Reader
}

type DvHttpResponse struct {
	Body           []byte
	Headers        http.Header
	StatusCode     int
	Location       *url.URL
	LocationParams url.Values
}

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

func NewClient(inputs *ClientInput) (*Client, error) {
	hostUrl := baseURL.ResolveReference(&url.URL{}).String()
	if inputs.HostURL != "" {
		hostUrl = inputs.HostURL
	}

	fmt.Printf("Using host: %v \n", hostUrl)
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, fmt.Errorf("Got error while creating cookie jar %s", err.Error())
	}
	c := Client{
		HTTPClient: &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
			Timeout: 10 * time.Second,
			Jar:     jar},
		HostURL: hostUrl,
	}

	if inputs.Username == "" || inputs.Password == "" {
		return nil, fmt.Errorf("User or Password not found")
	}

	c.Auth = AuthStruct{
		Username: inputs.Username,
		Password: inputs.Password,
	}

	// Use P1SSO if available
	if inputs.AuthP1SSO.PingOneAdminEnvId != "" || inputs.AuthP1SSO.PingOneTargetEnvId != "" {
		c.AuthP1SSO = inputs.AuthP1SSO
	}
	err = c.doSignIn()
	if err != nil {
		return nil, fmt.Errorf("Sign In failed with: %v", err)
	}

	return &c, nil
}

func (c *Client) doSignIn() error {
	if c.AuthP1SSO.PingOneAdminEnvId != "" || c.AuthP1SSO.PingOneTargetEnvId != "" {
		ar, err := c.SignInSSO()
		if err != nil {
			return err
		}
		c.Token = ar.AccessToken
		return nil
	}

	//Default Env User login
	ar, err := c.SignIn()
	if err != nil {
		return err
	}
	c.Token = ar.AccessToken
	return nil
}

func (c *Client) doRequestVerbose(req *http.Request, authToken *string, args *Params) (*DvHttpResponse, error) {
	token := c.Token
	// fmt.Printf("req is: %v", req)

	if authToken != nil {
		token = *authToken
		var bearer = "Bearer " + token
		req.Header.Add("Authorization", bearer)
	}
	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	}
	if args != nil {
		req.URL.RawQuery = args.QueryParams().Encode()
	}
	// fmt.Printf("client request is: %v", req)
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	rbody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusFound {
		return nil, fmt.Errorf("status: %d, rbody: %s", res.StatusCode, rbody)
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

	return &resp, err
}

func (c *Client) doRequest(req *http.Request, authToken *string, args *Params) ([]byte, *http.Response, error) {
	token := c.Token
	if authToken != nil {
		token = *authToken
	}
	var bearer = "Bearer " + token
	req.Header.Add("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	if args != nil {
		req.URL.RawQuery = args.QueryParams().Encode()
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}
	return body, res, err
}

func (c *Client) doRequestRetryable(req DvHttpRequest, authToken *string, args *Params) ([]byte, error) {
	// req.Close = true
	// fmt.Printf("req.body is: %v", req.Body)
	// urlRetry := fmt.Sprintf("%s://%s%s", req.URL.Scheme, req.URL.Host, req.URL.Path)
	// bodyRetry, err := req.Clone()
	// if err != nil {
	// 	return nil, err
	// }

	reqInit, err := http.NewRequest(req.Method, req.Url, req.Body)
	if err != nil {
		return nil, err
	}
	reqRetry, err := http.NewRequest(req.Method, req.Url, req.Body)
	if err != nil {
		return nil, err
	}
	body, res, err := c.doRequest(reqInit, authToken, args)
	if err != nil {
		return nil, err
	}
	if res.StatusCode == http.StatusUnauthorized && c.AuthRefresh == false {
		if err != nil {
			return nil, err
		}
		err = c.refreshAuth()
		if err != nil {
			return nil, err
		}
		_, err := c.SetEnvironment(&c.CompanyID)
		if err != nil {
			return nil, err
		}
		// fmt.Printf("req.body retry is: %v", reqRetry.Body)
		var resRetry *http.Response
		var bodyRetry []byte
		bodyRetry, resRetry, err = c.doRequest(reqRetry, authToken, args)
		if err != nil {
			fmt.Printf("error: %v", err)
			return nil, err
		}
		res = resRetry
		body = bodyRetry
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}
	return body, err
}

// refreshAuth is used to rerun the sign-on process.
// This is useful when the client's initial access_token was made before
// the target environment was created. (common in Terraform)
func (c *Client) refreshAuth() error {
	c.AuthRefresh = true
	// c.HTTPClient.Jar = nil
	// jar, err := cookiejar.New(nil)
	// if err != nil {
	// 	return fmt.Errorf("Got error while creating cookie jar %s", err.Error())
	// }
	// c.HTTPClient.Jar = jar
	err := c.doSignIn()
	if err != nil {
		return fmt.Errorf("Refreshing Sign In failed with: %v", err)
	}
	return nil
}
