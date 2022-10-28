package svc

type VerifyToken interface {
	Do(tkn string) (Payload, bool, error)
}

type Payload struct {
	UserID string
}
