package authn

import (
	"encoding/json"

	svcrouter "github.com/onyanko-pon/monorepo/server/svc_router"
)

type Token interface {
	VerifyToken(token string) (VerifyRes, error)
	EncodeToken(p Payload) (EncodeRes, error)
}
type TokenImple struct{}

type VerifyRes struct {
	UserID   string
	Verified bool
}

func (a TokenImple) VerifyToken(token string) (VerifyRes, error) {
	vrq := svcrouter.TokenVerifyReq{
		Token: token,
	}
	j, err := json.Marshal(vrq)
	if err != nil {
		return VerifyRes{}, err
	}
	vrs, err := svcrouter.Handle(svcrouter.UserTokenVerify, string(j))
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

type Payload struct {
	UserID string
}

type EncodeRes struct {
	Token string
}

func (a TokenImple) EncodeToken(p Payload) (EncodeRes, error) {
	vrq := svcrouter.TokenEncodeReq{
		UserID: p.UserID,
	}
	j, err := json.Marshal(vrq)
	if err != nil {
		return EncodeRes{}, err
	}
	vrs, err := svcrouter.Handle(svcrouter.UserTokenEncode, string(j))
	if err != nil {
		return EncodeRes{}, err
	}
	var rs svcrouter.TokenEncodeRes
	err = json.Unmarshal([]byte(vrs), &rs)
	if err != nil {
		return EncodeRes{}, err
	}
	return EncodeRes{
		Token: rs.Token,
	}, nil
}
