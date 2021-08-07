package gosnail

import (
	"gosnail/core"
	"net/http"
)

type Response struct {
	w           http.ResponseWriter
	statusCode  int
	headers     *core.Headers
	headersSent bool
}

func (r *Response) Headers() *core.Headers {
	return r.headers
}

func (r *Response) StatusCode() int {
	return r.statusCode
}

func (r *Response) SetStatusCode(code int) {
	r.statusCode = code
}

func (r *Response) Write(data []byte) (int, error) {
	return r.w.Write(data)
}

func (r *Response) End() error {
	_, err := r.Write([]byte{})
	return err
}
