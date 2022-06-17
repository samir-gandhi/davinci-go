package davinci

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

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

func (c *Client) doRequest(req *http.Request, authToken *string) ([]byte, error) {
	token := c.Token

	if authToken != nil {
		token = *authToken
	}
	var bearer = "Bearer " + token
	req.Header.Add("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

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
