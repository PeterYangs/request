package request

import (
	"net/http"
	"net/url"
)

type Client struct {
	client *http.Client
	header map[string]string
}

func NewClient() *Client {

	return &Client{
		client: &http.Client{},
	}
}

// Header 全局header
func (c *Client) Header(header map[string]string) *Client {

	c.header = header

	return c
}

func (c *Client) Proxy(proxyUrl string) *Client {

	c.client.Transport = &http.Transport{
		Proxy: func(r *http.Request) (*url.URL, error) {

			return url.Parse(proxyUrl)
		},
	}

	return c
}

func (c *Client) R() *request {

	return &request{client: c.client, header: c.header}
}
