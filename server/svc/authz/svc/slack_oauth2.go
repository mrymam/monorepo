package svc

import (
	"net/http"

	"github.com/slack-go/slack"
)

type SlackOAuth2Svc struct {
	clientID     string
	clientSecret string
	redirectURL  string
}

func (s SlackOAuth2Svc) GetAccessToken(authnCode string) (string, error) {
	clt := http.DefaultClient
	res, err := slack.GetOAuthV2Response(clt, s.clientID, s.clientSecret, authnCode, s.redirectURL)
	if err != nil {
		return "", err
	}
	return res.AccessToken, nil
}
