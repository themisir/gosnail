package gosnail

import (
	"io"
	"net/http"
	"net/url"

	"github.com/themisir/gosnail/core"
)

type Request struct {
	Method      string
	URL         *url.URL
	QueryParams *url.Values
	body        io.ReadCloser
	headers     *core.Headers
	req         *http.Request
}

func (r *Request) Headers() *core.Headers {
	return r.headers
}

func (r *Request) Query(name string) string {
	return r.QueryParams.Get(name)
}

func (r *Request) Body() io.ReadCloser {
	return r.body
}
