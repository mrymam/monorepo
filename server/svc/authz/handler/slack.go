package handler

import (
	adapter "github.com/mrymam/monorepo/server/adapter/svc/authz"
	"github.com/mrymam/monorepo/server/svc/authz/domain/svc"
	svcimpl "github.com/mrymam/monorepo/server/svc/authz/svc"
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

func (a SlackOAuth2) GetAccessToken(oauthCode adapter.SlackOAuthCode) (adapter.SlackAccessToken, error) {
	accessToken, err := a.svc.GetAccessToken(string(oauthCode))
	if err != nil {
		return "", err
	}
	return adapter.SlackAccessToken(accessToken), nil
}
