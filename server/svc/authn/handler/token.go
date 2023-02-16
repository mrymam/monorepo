package handler

import (
	adapter "github.com/onyanko-pon/monorepo/server/adapter/svc/authn"
	"github.com/onyanko-pon/monorepo/server/svc/authn/di"
	"github.com/onyanko-pon/monorepo/server/svc/authn/domain/model"
	"github.com/onyanko-pon/monorepo/server/svc/authn/domain/svc"
)

func InitAuthn() (Authn, error) {
	svc, err := di.GetTokenSvc()
	if err != nil {
		return Authn{}, nil
	}
	return Authn{
		tokenSvc: svc,
	}, nil
}

type Authn struct {
	tokenSvc svc.TokenSvc
}

func (a Authn) VerifyToken(token adapter.Token) (adapter.UserID, adapter.Varid, error) {
	payload, err := a.tokenSvc.Parse(string(token))
	return adapter.UserID(payload.UserID), err == nil, err
}

func (a Authn) EncodeToken(userid adapter.UserID) (adapter.Token, error) {
	payload := svc.Payload{
		UserID: model.UserID(userid),
	}
	token, err := a.tokenSvc.Encode(payload)
	if err != nil {
		return "", err
	}
	return adapter.Token(token), nil
}
