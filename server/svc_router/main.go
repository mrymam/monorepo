package svcrouter

import "fmt"

var handlers []Handler

func init() {
	handlers = []Handler{}
}

func AddHandler(key Key, f Func) error {
	h := Handler{Key: key, Func: f}
	handlers = append(handlers, h)
	return nil
}

func Handle(key Key, body string) (string, error) {
	for _, h := range handlers {
		if h.Key == key {
			return h.Func(body)
		}
	}
	return "", fmt.Errorf("handler not found: key: %s", key)
}
