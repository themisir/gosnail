package main

import (
	"io"
	"net/http"
	"net/url"
)

type Request struct {
	Method      string
	URL         *url.URL
	QueryParams *url.Values
	body        io.ReadCloser
	headers     *Headers
	req         *http.Request
}

func (r *Request) Headers() *Headers {
	return r.headers
}

func (r *Request) Query(name string) string {
	return r.QueryParams.Get(name)
}

func (r *Request) Body() io.ReadCloser {
	return r.body
}
