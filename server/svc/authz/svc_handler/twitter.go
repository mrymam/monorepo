package svchandler

import (
	"encoding/json"

	"github.com/onyanko-pon/monorepo/server/svc/authz/domain/svc"
	"github.com/onyanko-pon/monorepo/server/svc/authz/infra/twitter"
	svcimpl "github.com/onyanko-pon/monorepo/server/svc/authz/svc"
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

func (a TwitterAuth1) VerifyAccessToken(arg string) (string, error) {
	var req svcrouter.TwitterOAuth1VerifyAccessTokenReq
	err := json.Unmarshal([]byte(arg), &req)
	if err != nil {
		return "", err
	}
	ujson, err := a.svc.VerifyAccessToken(svc.AccessToken(req.AccessToken), svc.AccessSecret(req.AccessSecret))
	if err != nil {
		return "", err
	}
	var u twitter.User
	if err := json.Unmarshal([]byte(ujson), &u); err != nil {
		return "", err
	}

	rs := svcrouter.TwitterOAuth1VerifyAccessTokenRes{
		User: svcrouter.TwitterUser{
			ID:              u.ID,
			Name:            u.Name,
			ScreenName:      u.ScreenName,
			ProfileImageUrl: u.ProfileImageUrl,
		},
	}
	j, err := json.Marshal(rs)
	return string(j), err
}
