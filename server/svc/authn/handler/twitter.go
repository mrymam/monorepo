package handler

import (
	"github.com/dghubble/oauth1"
	adapter "github.com/onyanko-pon/monorepo/server/adapter/svc/authn"
	"github.com/onyanko-pon/monorepo/server/svc/authn/domain/model"
	"github.com/onyanko-pon/monorepo/server/svc/authn/infra/repository"
)

type TwitterAuth struct {
	twrepo   repository.TwitterUserRepo
	authrepo repository.TwitterUserIdentityRepo
}

func InitTwitterAuth() (TwitterAuth, error) {
	twrepo, err := repository.InitTwitterUserRepo()
	if err != nil {
		return TwitterAuth{}, err
	}
	authrepo, err := repository.InitTwitterUserIdentityRepo()
	if err != nil {
		return TwitterAuth{}, err
	}
	return TwitterAuth{
		twrepo:   twrepo,
		authrepo: authrepo,
	}, nil
}

func (a TwitterAuth) Authenticate(accessToken adapter.TwitterAccessToken, accessSecret adapter.TwitterAccessSecret) (adapter.UserID, adapter.TwitterProfile, error) {
	token := oauth1.NewToken(string(accessToken), string(accessSecret))
	u, err := a.twrepo.Get(token)
	if err != nil {
		return "", adapter.TwitterProfile{}, err
	}
	ex, err := a.authrepo.Exist(u.ID)
	if err != nil {
		return "", adapter.TwitterProfile{}, err
	}
	if !ex {
		m, err := model.InitTwitterUserIdentity(u.ID)
		if err != nil {
			return "", adapter.TwitterProfile{}, err
		}
		_, err = a.authrepo.Create(m)
		if err != nil {
			return "", adapter.TwitterProfile{}, err
		}
	}
	at, err := a.authrepo.GetByTiwtterUserID(u.ID)
	if err != nil {
		return "", adapter.TwitterProfile{}, err
	}

	return adapter.UserID(at.TwitterUserID), adapter.TwitterProfile{
		ID:         string(u.ID),
		ScreenName: string(u.Name),
		Name:       string(u.Name),
		ImageURL:   string(u.ProfileImageUrl),
	}, nil
}
