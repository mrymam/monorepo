package svc

import (
	"github.com/onyanko-pon/monorepo/server/svc/authn/domain/model"
)

type VerifyToken interface {
	Do(tkn string) (Payload, bool, error)
}

type Payload struct {
	UserID model.UserID
}
