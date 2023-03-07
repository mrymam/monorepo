package handler

import (
	"github.com/dghubble/oauth1"
	adapter "github.com/mrymam/radio_rec/server/adapter/svc/authn"
	"github.com/mrymam/radio_rec/server/svc/authn/domain/model"
	"github.com/mrymam/radio_rec/server/svc/authn/infra/repository"
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

func (a TwitterAuth) Authenticate(req adapter.TwitterAuthenticateReq) (adapter.TwitterAuthenticateRes, error) {
	token := oauth1.NewToken(req.AccessToken, req.AccessSecret)
	u, err := a.twrepo.Get(token)
	if err != nil {
		return adapter.TwitterAuthenticateRes{}, err
	}
	ex, err := a.authrepo.Exist(u.ID)
	if err != nil {
		return adapter.TwitterAuthenticateRes{}, err
	}
	if !ex {
		m, err := model.InitTwitterUserIdentity(u.ID)
		if err != nil {
			return adapter.TwitterAuthenticateRes{}, err
		}
		_, err = a.authrepo.Create(m)
		if err != nil {
			return adapter.TwitterAuthenticateRes{}, err
		}
	}
	at, err := a.authrepo.GetByTiwtterUserID(u.ID)
	if err != nil {
		return adapter.TwitterAuthenticateRes{}, err
	}

	return adapter.TwitterAuthenticateRes{
		UserID: string(at.UserID),
		Profile: adapter.TwitterProfile{
			ID:         string(u.ID),
			ScreenName: string(u.Name),
			Name:       string(u.Name),
			ImageURL:   string(u.ProfileImageUrl),
		},
	}, nil
}
