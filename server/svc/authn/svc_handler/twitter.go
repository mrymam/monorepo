package svchandler

import (
	"encoding/json"

	"github.com/onyanko-pon/monorepo/server/svc/authn/domain/svc"
	svcimpl "github.com/onyanko-pon/monorepo/server/svc/authn/svc"
	svcrouter "github.com/onyanko-pon/monorepo/server/svc_router"
)

type TwitterAuth1 struct {
	svc svc.OAuth1Svc
}

func InitTwitterAuth1() (TwitterAuth1, error) {
	s, err := svcimpl.InitTwitterAuth1Svc()
	if err != nil {
		return TwitterAuth1{}, nil
	}
	return TwitterAuth1{
		svc: s,
	}, nil
}

func (a TwitterAuth1) GetAccessToken(arg string) (string, error) {
	var req svcrouter.TwitterOAuth1GetAccessTokenReq
	err := json.Unmarshal([]byte(arg), &req)
	if err != nil {
		return "", err
	}
	accessToken, accessSecret, err := a.svc.FetchAccessToken(
		svc.OAuthToken(req.OAuthToken),
		svc.OAuthSecret(req.OAuthSecret),
		svc.OAuthVerifier(req.OAuthVerifier))
	if err != nil {
		return "", err
	}

	rs := svcrouter.TwitterOAuth1GetAccessTokenRes{
		AccessToken:  string(accessToken),
		AccessSecret: string(accessSecret),
	}
	j, err := json.Marshal(rs)
	return string(j), err
}
