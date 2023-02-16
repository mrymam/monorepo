package handler

import (
	adapter "github.com/mrymam/monorepo/server/adapter/svc/authn"
	"github.com/mrymam/monorepo/server/svc/authn/di"
	"github.com/mrymam/monorepo/server/svc/authn/domain/model"
	"github.com/mrymam/monorepo/server/svc/authn/domain/svc"
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

func (a Authn) VerifyToken(req adapter.AuthnVerifyTokenReq) (adapter.AuthnVerifyTokenRes, error) {
	payload, err := a.tokenSvc.Parse(req.Token)
	if err != nil {
		return adapter.AuthnVerifyTokenRes{}, err
	}
	return adapter.AuthnVerifyTokenRes{
		UserID: string(payload.UserID),
		Varid:  err == nil,
	}, nil
}

func (a Authn) EncodeToken(req adapter.AuthnEncodeTokenReq) (adapter.AuthnEncodeTokenRes, error) {
	payload := svc.Payload{
		UserID: model.UserID(req.UserID),
	}
	token, err := a.tokenSvc.Encode(payload)
	if err != nil {
		return adapter.AuthnEncodeTokenRes{}, err
	}
	return adapter.AuthnEncodeTokenRes{Token: token}, nil
}
