package authn

import (
	"encoding/json"

	svcrouter "github.com/onyanko-pon/monorepo/server/svc_router"
)

type Authn interface {
	Verify(token string) (VerifyRes, error)
}
type AuthnImpl struct{}

type VerifyRes struct {
	UserID   string
	Verified bool
}

func (a AuthnImpl) Verify(token string) (VerifyRes, error) {
	vrq := svcrouter.TokenVerifyReq{
		Token: token,
	}
	j, err := json.Marshal(vrq)
	if err != nil {
		return VerifyRes{}, err
	}
	vrs, err := svcrouter.Handle(svcrouter.UserVerify, string(j))
	if err != nil {
		return VerifyRes{}, err
	}
	var rs svcrouter.TokenVerifyRes
	err = json.Unmarshal([]byte(vrs), &rs)
	if err != nil {
		return VerifyRes{}, err
	}
	return VerifyRes{
		UserID:   rs.UserID,
		Verified: rs.Verified,
	}, nil
}
