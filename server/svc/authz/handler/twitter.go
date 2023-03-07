package handler

import (
	adapter "github.com/mrymam/radio_rec/server/adapter/svc/authz"
	"github.com/mrymam/radio_rec/server/svc/authz/domain/svc"
	svcimpl "github.com/mrymam/radio_rec/server/svc/authz/svc"
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

func (a TwitterAuth1) GetAccessToken(
	token adapter.TwitterOAuthToken,
	secret adapter.TwitterOAuthSecret,
	verifier adapter.TwitterOAuthVerifier) (
	adapter.TwitterAccessToken,
	adapter.TwitterAccessSecret,
	error) {

	accessToken, accessSecret, err := a.svc.FetchAccessToken(
		svc.OAuthToken(token),
		svc.OAuthSecret(secret),
		svc.OAuthVerifier(verifier))
	if err != nil {
		return "", "", err
	}

	return adapter.TwitterAccessToken(accessToken), adapter.TwitterAccessSecret(accessSecret), err
}
