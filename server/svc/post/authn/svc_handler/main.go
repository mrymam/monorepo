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
	res := svcrouter.UserVerifyRes{
		UserID:   "xxxxx.xxxxx",
		Verified: true,
	}
	j, err := json.Marshal(res)
	return string(j), err
}
