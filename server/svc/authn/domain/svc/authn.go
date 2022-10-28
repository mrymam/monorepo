package svc

import (
	"github.com/onyanko-pon/monorepo/server/svc/authn/domain/model"
)

type TokenSvc interface {
	Parse(tkn string) (Payload, error)
	Encode(Payload) (string, error)
}

type Payload struct {
	UserID model.UserID
}
