package svchandler

import (
	"encoding/json"

	"github.com/onyanko-pon/monorepo/server/svc/authn/domain/svc"
	svcrouter "github.com/onyanko-pon/monorepo/server/svc_router"
)

type Authn struct {
	tokenSvc svc.TokenSvc
}

func (a Authn) Verify(arg string) (string, error) {
	var req svcrouter.UserVerifyReq
	err := json.Unmarshal([]byte(arg), &req)
	if err != nil {
		return "", err
	}

	payload, err := a.tokenSvc.Parse(req.Token)
	rs := svcrouter.UserVerifyRes{
		UserID:   string(payload.UserID),
		Verified: err == nil,
	}
	j, err := json.Marshal(rs)
	return string(j), err
}
