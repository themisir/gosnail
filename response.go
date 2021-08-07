package main

import (
	"errors"
	"net/http"
)

type Response struct {
	StatusCode  int
	w           http.ResponseWriter
	headers     *Headers
	headersSent bool
}

func (r *Response) Headers() *Headers {
	return r.headers
}

func (r *Response) SendHeaders() error {
	if r.headersSent {
		return errors.New("headers already sent")
	}

	for k, v := range r.headers.values {
		r.w.Header().Add(k, v)
	}

	r.w.WriteHeader(r.StatusCode)
	r.headersSent = true

	return nil
}

func (r *Response) Write(data []byte) (int, error) {
	if !r.headersSent {
		r.SendHeaders()
	}
	return r.w.Write(data)
}

func (r *Response) End() error {
	_, err := r.Write([]byte{})
	return err
}
