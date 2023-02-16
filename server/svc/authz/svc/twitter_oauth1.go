package svc

import (
	"github.com/dghubble/oauth1"
	"github.com/dghubble/oauth1/twitter"
	"github.com/mrymam/monorepo/server/pkg/setting"
	"github.com/mrymam/monorepo/server/svc/authz/domain/svc"
)

type TwitterAuth1Svc struct {
	cfg oauth1.Config
}

func InitTwitterAuth1Svc() (TwitterAuth1Svc, error) {
	s := setting.Get().Twitter.OAuth1
	return TwitterAuth1Svc{
		cfg: oauth1.Config{
			ConsumerKey:    s.ConsumerKey,
			ConsumerSecret: s.ConsumerSecret,
			CallbackURL:    s.CallbackURL,
			Endpoint:       twitter.AuthorizeEndpoint,
		},
	}, nil
}

func (s TwitterAuth1Svc) FetchAccessToken(oauthToken svc.OAuthToken, oauthSecret svc.OAuthSecret, oauthVerifier svc.OAuthVerifier) (svc.AccessToken, svc.AccessSecret, error) {
	accessToken, accessSecret, err := s.cfg.AccessToken(string(oauthToken), string(oauthSecret), string(oauthVerifier))
	return svc.AccessToken(accessToken), svc.AccessSecret(accessSecret), err
}
