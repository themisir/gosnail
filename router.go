package main

import (
	"net/http"
	"strings"
)

type RouterHandler = func(*Context, func())

const Any = ""

type RouterEntry struct {
	method    string
	path      string
	pathRange bool
	handlers  []RouterHandler
	next      *RouterEntry
}

type Router struct {
	head *RouterEntry
	foot *RouterEntry
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Next(ctx *Context, next func()) {
	for entry := r.head; entry != nil; entry = entry.next {

		if len(entry.handlers) == 0 {
			continue
		}

		if entry.method != "" && entry.method != ctx.req.Method {
			continue
		}

		if entry.path != "" && entry.path != ctx.req.URL.Path {
			if entry.pathRange {
				match := entry.path

				if !strings.HasSuffix(match, "/") {
					match = match + "/"
				}

				if !strings.HasPrefix(ctx.req.URL.Path, match) {
					continue
				}
			} else {
				continue
			}
		}

		doContinue := false

		for _, handler := range entry.handlers {
			doContinue = false
			handler(ctx, func() {
				doContinue = true
			})

			if !doContinue {
				break
			}
		}

		if !doContinue {
			break
		}
	}

	next()
}

func (r *Router) Handle(method string, path string, handlers ...RouterHandler) {
	if len(handlers) > 0 {
		entry := &RouterEntry{
			path:     path,
			method:   method,
			handlers: handlers,
		}

		if r.foot == nil {
			r.head = entry
			r.foot = entry
		} else {
			r.foot.next = entry
			r.foot = entry
		}
	}
}

func (r *Router) Use(path string, handlers ...RouterHandler) {
	r.Handle("", path, handlers...)
	r.foot.pathRange = true
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
