package svchandler

import (
	"encoding/json"

	"github.com/onyanko-pon/monorepo/server/svc/authz/domain/svc"
	svcimpl "github.com/onyanko-pon/monorepo/server/svc/authz/svc"
	svcrouter "github.com/onyanko-pon/monorepo/server/svc_router"
)

type SlackOAuth2 struct {
	svc svc.OAuth2Svc
}

func InitSlackOAuth2() (SlackOAuth2, error) {
	s, err := svcimpl.InitSlackAuth2Svc()
	if err != nil {
		return SlackOAuth2{}, nil
	}
	return SlackOAuth2{
		svc: s,
	}, nil
}

func (a SlackOAuth2) GetAccessToken(arg string) (string, error) {
	var req svcrouter.SlackOAuth2GetAccessTokenReq
	err := json.Unmarshal([]byte(arg), &req)
	if err != nil {
		return "", err
	}
	accessToken, err := a.svc.GetAccessToken(req.OAuthCode)
	if err != nil {
		return "", err
	}

	rs := svcrouter.SlackOAuth2GetAccessTokenRes{
		AccessToken: string(accessToken),
	}
	j, err := json.Marshal(rs)
	return string(j), err
}
