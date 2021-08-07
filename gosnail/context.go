package gosnail

type Context struct {
	req    *Request
	res    *Response
	values map[string]interface{}
}

func (c *Context) Request() *Request {
	return c.req
}

func (c *Context) Response() *Response {
	return c.res
}

func (c *Context) Get(key string) interface{} {
	return c.values[key]
}

func (c *Context) Set(key string, value interface{}) {
	c.values[key] = value
}

func (c *Context) Delete(key string) {
	delete(c.values, key)
}
