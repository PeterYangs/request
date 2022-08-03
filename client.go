package request

import (
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	client     *http.Client
	header     map[string]string
	timeout    time.Duration
	retryTimes int
	debug      bool
	transport  *http.Transport
}

func NewClient() *Client {

	return &Client{
		client:    &http.Client{},
		transport: http.DefaultTransport.(*http.Transport).Clone(),
	}
}

// Header 全局header
func (c *Client) Header(header map[string]string) *Client {

	c.header = header

	return c
}

func (c *Client) Proxy(proxyUrl string) *Client {

	//c.client.Transport = &http.Transport{
	//	Proxy: func(r *http.Request) (*url.URL, error) {
	//
	//		return url.Parse(proxyUrl)
	//	},
	//	MaxConnsPerHost: 10,
	//}

	c.transport.Proxy = func(r *http.Request) (*url.URL, error) {

		return url.Parse(proxyUrl)
	}

	c.transport.MaxConnsPerHost = 10

	return c
}

func (c *Client) Timeout(timeout time.Duration) *Client {

	c.timeout = timeout

	return c

}

func (c *Client) ReTry(times int) *Client {

	c.retryTimes = times

	return c

}

func (c *Client) Debug() *Client {

	c.debug = true

	return c

}

func (c *Client) GetTransport() *http.Transport {

	return c.transport
}

func (c *Client) R() *request {

	c.client.Transport = c.transport

	return newRequest(c)
}
