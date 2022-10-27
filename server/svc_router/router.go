package svcrouter

type Handler struct {
	Key  Key
	Func Func
}

type Func func(string) (string, error)
