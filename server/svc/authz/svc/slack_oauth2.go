package svc

import (
	"net/http"

	"github.com/mrymam/monorepo/server/pkg/setting"
	"github.com/slack-go/slack"
)

type SlackOAuth2Svc struct {
	clientID     string
	clientSecret string
	redirectURL  string
}

func InitSlackAuth2Svc() (SlackOAuth2Svc, error) {
	s := setting.Get().Slack.OAuth2
	return SlackOAuth2Svc{
		clientID:     s.ClientID,
		clientSecret: s.ClientSecret,
		redirectURL:  s.RedirectURL,
	}, nil
}

func (s SlackOAuth2Svc) GetAccessToken(authnCode string) (string, error) {
	clt := http.DefaultClient
	res, err := slack.GetOAuthV2Response(clt, s.clientID, s.clientSecret, authnCode, s.redirectURL)
	if err != nil {
		return "", err
	}
	return res.AccessToken, nil
}
