package request

import (
	"net/http"
)

type response struct {
	response *http.Response
	request  *request
}

func NewResponse(rsp *http.Response, req *request) *response {

	return &response{rsp, req}
}

func (r *response) Header() http.Header {

	return r.response.Header
}

func (r *response) Content() (*content, error) {

	return NewContent(r)
}
