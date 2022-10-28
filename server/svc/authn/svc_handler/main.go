package svchandler

import (
	"encoding/json"

	svcrouter "github.com/onyanko-pon/monorepo/server/svc_router"
)

type Authn struct{}

func (a Authn) Verify(arg string) (string, error) {
	var req svcrouter.UserVerifyReq
	err := json.Unmarshal([]byte(arg), &req)
	if err != nil {
		return "", err
	}

	// TODO: implement jwt verify
	rs := svcrouter.UserVerifyRes{
		UserID:   "xxxxx.xxxxx",
		Verified: false,
	}
	j, err := json.Marshal(rs)
	return string(j), err
}
