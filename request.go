package request

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"runtime/debug"
)

type Client struct {
	client *http.Client
}

type request struct {
	request *http.Response
	client  *http.Client
}

type response struct {
	response *http.Response
}

type body struct {
	body io.ReadCloser
}

type content struct {
	content []byte
}

func NewClient() *Client {

	return &Client{
		client: &http.Client{},
	}
}

func (c *Client) R() *request {

	return &request{client: c.client}
}

func (r *request) Get(url string) (*response, error) {

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {

		return nil, err
	}

	rsp, err := r.client.Do(req)

	return &response{response: rsp}, nil

}

func (r *response) Body() body {

	return body{
		body: r.response.Body,
	}
}

func (r *response) Header() http.Header {

	return r.response.Header
}

func (b body) Content() content {

	bb, err := ioutil.ReadAll(b.body)

	if err != nil {

		log.Println(err)
		log.Println(debug.Stack())

		return content{content: []byte{}}
	}

	defer b.body.Close()

	return content{content: bb}
}

func (c content) ToString() string {

	return string(c.content)
}
