package gosnail

import (
	"gosnail/core"
	"net/http"
)

type Handler interface {
	Next(ctx *Context, next func())
}

type Application struct {
	handler Handler
}

func NewApplication(handler Handler) *Application {
	return &Application{handler}
}

func (a *Application) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	req := &Request{
		URL:         r.URL,
		Method:      r.Method,
		QueryParams: &query,
		body:        r.Body,
		headers:     core.NewHeadersFromValues(r.Header),
		req:         r,
	}
	res := &Response{
		w:           w,
		statusCode:  http.StatusOK,
		headers:     core.NewHeaders(),
		headersSent: false,
	}
	ctx := &Context{
		req:    req,
		res:    res,
		values: make(map[string]interface{}),
	}

	a.handler.Next(ctx, func() {
		ctx.Response().SetStatusCode(404)
		ctx.Response().End()
	})
}

func (a *Application) Listen(addr string) error {
	return http.ListenAndServe(addr, a)
}

func (a *Application) ListenTLS(addr, certFile, keyFile string) error {
	return http.ListenAndServeTLS(addr, certFile, keyFile, a)
}