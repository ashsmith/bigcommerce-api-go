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

type Client struct {
	app        App
	HTTPClient *http.Client
	Webhooks   WebhooksService
}

func (a App) NewClient(httpClient *http.Client) *Client {
	return NewClient(a, httpClient)
}

func NewClient(app App, httpClient *http.Client) *Client {
	c := &Client{
		app: app,
	}

	c.Webhooks = &WebhooksServiceOp{client: c}

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

	return body, nil
}
