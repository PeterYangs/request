package request

import (
	"net/http"
)

type response struct {
	response *http.Response
}

func (r *response) Body() body {

	return body{
		body: r.response.Body,
	}
}

func (r *response) Header() http.Header {

	return r.response.Header
}
