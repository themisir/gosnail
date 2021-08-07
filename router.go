package gosnail

import "net/http"

type RouterHandler = func(*Context, func())

type RouterEntry struct {
	method   string
	path     string
	handlers []RouterHandler
}

type Router struct {
	entries []*RouterEntry
}

func NewRouter() *Router {
	return &Router{
		entries: []*RouterEntry{},
	}
}

func (r *Router) Next(ctx *Context, next func()) {
	for _, entry := range r.entries {
		if len(entry.handlers) == 0 {
			continue
		}

		if ctx.req.Method != entry.method {
			continue
		}

		if ctx.req.URL.Path != entry.path {
			continue
		}

		if len(entry.handlers) == 1 {
			entry.handlers[0](ctx, next)
			return
		}

		var nextFunc func()
		i := -1

		nextFunc = func() {
			i++
			if i < len(entry.handlers) {
				entry.handlers[i](ctx, nextFunc)
			} else {
				next()
			}
		}

		nextFunc()
		return
	}

	next()
}

func (r *Router) Handle(method string, path string, handlers ...RouterHandler) {
	if len(handlers) > 0 {
		r.entries = append(r.entries, &RouterEntry{
			path:     path,
			method:   method,
			handlers: handlers,
		})
	}
}

func (r *Router) Get(path string, handlers ...RouterHandler) {
	r.Handle(http.MethodGet, path, handlers...)
}

func (r *Router) Head(path string, handlers ...RouterHandler) {
	r.Handle(http.MethodHead, path, handlers...)
}

func (r *Router) Post(path string, handlers ...RouterHandler) {
	r.Handle(http.MethodPost, path, handlers...)
}

func (r *Router) Put(path string, handlers ...RouterHandler) {
	r.Handle(http.MethodPut, path, handlers...)
}

func (r *Router) Patch(path string, handlers ...RouterHandler) {
	r.Handle(http.MethodPatch, path, handlers...)
}

func (r *Router) Delete(path string, handlers ...RouterHandler) {
	r.Handle(http.MethodDelete, path, handlers...)
}

func (r *Router) Connect(path string, handlers ...RouterHandler) {
	r.Handle(http.MethodConnect, path, handlers...)
}

func (r *Router) Options(path string, handlers ...RouterHandler) {
	r.Handle(http.MethodOptions, path, handlers...)
}

func (r *Router) Trace(path string, handlers ...RouterHandler) {
	r.Handle(http.MethodTrace, path, handlers...)
}
