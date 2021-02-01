package bigcommerce

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// App represents basic app settings
type App struct {
	StoreHash   string
	ClientID    string
	AccessToken string
}

// Client is a collection of services that interacts with the BigCommerce API.
type Client struct {
	app        App
	HTTPClient http.Client
	Webhooks   WebhooksService
	Storefront StorefrontService
}

type Links struct {
	Previous string `json:"previous,omitempty"`
	Next     string `json:"next,omitempty"`
	Current  string `json:"current,omitempty"`
}

type PaginationResult struct {
	Total       int64 `json:"total,omitempty"`
	Count       int64 `json:"count,omitempty"`
	PerPage     int64 `json:"per_page,omitempty"`
	CurrentPage int64 `json:"current_page,omitempty"`
	Totalpages  int64 `json:"total_pages,omitempty`
	Links       Links `json:"links,omitempty"`
}

type MetaResult struct {
	Pagination PaginationResult `json:"pagination,omitempty"`
}

// NewClient will create a new client instance for interacting with the BigCommerce API.
func (a App) NewClient(httpClient http.Client) *Client {
	return client(a, httpClient)
}

func client(app App, httpClient http.Client) *Client {
	c := &Client{
		app: app,
	}

	c.Webhooks = &WebhooksServiceOp{client: c}

	c.Storefront = StorefrontService{}
	c.Storefront.Status = &StorefrontStatusOp{client: c}
	c.Storefront.Seo = &StorefrontSeoSettingsOp{client: c}
	c.Storefront.Security = &StorefrontSecuritySettingsOp{client: c}
	c.Storefront.Search = &StorefrontSearchSettingsOp{client: c}
	c.Storefront.Category = &StorefrontCategorySettingsOp{client: c}
	c.Storefront.RobotsTxt = &StorefrontRobotsTxtSettingsOp{client: c}

	return c
}

func (c *Client) newRequest(method, path string, body io.Reader) (*http.Request, error) {
	url := fmt.Sprintf("https://api.bigcommerce.com/stores/%s%s", c.app.StoreHash, path)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return &http.Request{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Auth-Client", c.app.ClientID)
	req.Header.Set("X-Auth-Token", c.app.AccessToken)

	return req, nil
}

// DoRequest will create a request and return the response.
func (c *Client) DoRequest(method, path string, reqBody io.Reader) ([]byte, error) {
	req, err := c.newRequest(method, path, reqBody)
	if err != nil {
		return nil, err
	}

	res, doErr := c.HTTPClient.Do(req)
	if doErr != nil {
		return nil, doErr
	}

	defer res.Body.Close()

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return nil, readErr
	}

	if res.StatusCode >= 300 {
		return nil, fmt.Errorf("Received non-2xx status\nstatus code: %d\nbody: %s", res.StatusCode, string(body))
	}

	return body, nil
}
