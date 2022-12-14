package svchandler

import (
	"encoding/json"

	"github.com/onyanko-pon/monorepo/server/svc/authn/di"
	"github.com/onyanko-pon/monorepo/server/svc/authn/domain/model"
	"github.com/onyanko-pon/monorepo/server/svc/authn/domain/svc"
	svcrouter "github.com/onyanko-pon/monorepo/server/svc_router"
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

func (a Authn) VerifyToken(arg string) (string, error) {
	var req svcrouter.TokenVerifyReq
	err := json.Unmarshal([]byte(arg), &req)
	if err != nil {
		return "", err
	}

	payload, err := a.tokenSvc.Parse(req.Token)
	rs := svcrouter.TokenVerifyRes{
		UserID:   string(payload.UserID),
		Verified: err == nil,
	}
	j, err := json.Marshal(rs)
	return string(j), err
}

func (a Authn) EncodeToken(arg string) (string, error) {
	var req svcrouter.TokenEncodeReq
	err := json.Unmarshal([]byte(arg), &req)
	if err != nil {
		return "", err
	}
	payload := svc.Payload{
		UserID: model.UserID(req.UserID),
	}
	token, err := a.tokenSvc.Encode(payload)
	if err != nil {
		return "", err
	}
	rs := svcrouter.TokenEncodeRes{
		Token: token,
	}
	j, err := json.Marshal(rs)
	return string(j), err
}
