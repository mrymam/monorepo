package svc

import (
	"io/ioutil"

	"github.com/dghubble/oauth1"
	"github.com/dghubble/oauth1/twitter"
	"github.com/onyanko-pon/monorepo/server/pkg/setting"
	"github.com/onyanko-pon/monorepo/server/svc/authn/domain/svc"
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

func (s TwitterAuth1Svc) VerifyAccessToken(accessToken svc.AccessToken, accessSecret svc.AccessSecret) (string, error) {
	token := oauth1.NewToken(string(accessToken), string(accessSecret))
	clt := s.cfg.Client(oauth1.NoContext, token)

	path := "https://api.twitter.com/1.1/account/verify_credentials.json"
	resp, err := clt.Get(path)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
