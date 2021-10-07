package request

import (
	"net/http"
)

type response struct {
	response *http.Response
	request  *request
}

func (r *response) Body() body {

	return body{
		body:    r.response.Body,
		header:  r.response.Header,
		request: r.request,
	}
}

func (r *response) Header() http.Header {

	return r.response.Header
}
