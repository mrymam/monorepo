package svcrouter

import "fmt"

type Handler struct {
	Key  Key
	Func Func
}

type Func func(string) (string, error)

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
