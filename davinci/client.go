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
	HostURL    string
	HTTPClient *http.Client
	Token      string
	Auth       AuthStruct
	CompanyID  string
	AuthP1SSO  AuthP1SSO
}

type Params struct {
	Page        string
	Limit       string
	ExtraParams map[string]string
	// TODO: figure out what query is
	// query  string
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
		ar, err := c.SignInSSO()
		if err != nil {
			return nil, err
		}
		c.Token = ar.AccessToken
		c.CompanyID = ar.SelectedCompany
		return &c, nil
	}

	//Default Env User login
	ar, err := c.SignIn()
	if err != nil {
		return nil, err
	}
	c.Token = ar.AccessToken
	c.CompanyID = ar.SelectedCompany

	return &c, nil
}

func (c *Client) doRequestVerbose(req *http.Request, authToken *string, args *Params) (*DvHttpResponse, error) {
	token := c.Token

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
	fmt.Printf("client request is: %v", req)
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

func (c *Client) doRequest(req *http.Request, authToken *string, args *Params) ([]byte, error) {
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
	fmt.Printf("client request is: %v", req)
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
