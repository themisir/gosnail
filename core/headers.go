package core

type Headers struct {
	values map[string]string
}

func NewHeaders() *Headers {
	return &Headers{
		values: make(map[string]string),
	}
}

func NewHeadersFromValues(values map[string][]string) *Headers {
	headers := make(map[string]string)

	for k, v := range values {
		if len(v) > 0 {
			headers[k] = v[0]
		}
	}

	return &Headers{
		values: headers,
	}
}

func (h *Headers) Has(name string) bool {
	_, ok := h.values[name]
	return ok
}

func (h *Headers) Set(name string, value string) {
	h.values[name] = value
}

func (h *Headers) Get(name string) string {
	return h.values[name]
}

func (h *Headers) Delete(name string) {
	delete(h.values, name)
}
