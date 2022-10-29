package svchandler

import (
	"github.com/dghubble/oauth1"
	"github.com/onyanko-pon/monorepo/server/svc/authn/infra/repository"
)

type TwitterAuth struct {
	repo repository.TwitterUserRepo
}

func InitTwitterAuth() (TwitterAuth, error) {
	repo, err := repository.InitTwitterUserRepo()
	if err != nil {
		return TwitterAuth{}, err
	}
	return TwitterAuth{
		repo: repo,
	}, nil
}

func (a TwitterAuth) Authenticate(accessToken string, accessSecret string) (string, error) {
	token := oauth1.NewToken(accessToken, accessSecret)
	u, err := a.repo.Get(token)
	if err != nil {
		return "", err
	}
	return "", nil
}
