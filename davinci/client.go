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
	"net/url"
	"time"
)

// baseURL
var baseURL = url.URL{
	Scheme: "https",
	Host:   "api.singularkey.com",
	Path:   "/v1",
}

// const HostURL string = "https://api.singularkey.com"

type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
	Auth       AuthStruct
	CompanyID  string
}

type Params struct {
	Page  string
	Limit string
	// TODO: figure out what query is
	// query  string
}

func (args Params) QueryParams() url.Values {
	q := make(url.Values)

	if args.Page != "" {
		q.Add("page", args.Page)
	}

	if args.Limit != "" {
		q.Add("limit", args.Limit)
	}

	return q
}

func NewClient(host, username, password *string) (*Client, error) {
	if host != nil {
		baseURL.Host = *host
	}
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    baseURL.ResolveReference(&url.URL{}).String(),
	}

	if username == nil || password == nil {
		return nil, fmt.Errorf("User or Password not found")
	}

	c.Auth = AuthStruct{
		Username: *username,
		Password: *password,
	}

	ar, err := c.SignIn()
	if err != nil {
		return nil, err
	}
	c.Token = ar.AccessToken
	c.CompanyID = ar.SelectedCompany

	return &c, nil
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
