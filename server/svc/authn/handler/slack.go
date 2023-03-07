package handler

import (
	adapter "github.com/mrymam/radio_rec/server/adapter/svc/authn"
	"github.com/mrymam/radio_rec/server/svc/authn/domain/model"
	"github.com/mrymam/radio_rec/server/svc/authn/infra/repository"
	"github.com/slack-go/slack"
)

type SlackAuth struct {
	authrepo repository.SlackUserIdentyRepo
}

func InitSlackAuth() (SlackAuth, error) {
	authrepo, err := repository.InitSlackUserIdentityRepo()
	if err != nil {
		return SlackAuth{}, err
	}
	return SlackAuth{
		authrepo: authrepo,
	}, nil
}

func (a SlackAuth) Authenticate(param adapter.SlackAuthenticateReq) (adapter.SlackAuthenticateRes, error) {
	api := slack.New(param.AccessToken)
	ui, err := api.GetUserIdentity()
	if err != nil {
		return adapter.SlackAuthenticateRes{}, err
	}
	ex, err := a.authrepo.Exist(model.SlackUserID(ui.User.ID))
	if err != nil {
		return adapter.SlackAuthenticateRes{}, err
	}
	if !ex {
		m, err := model.InitTwitterSlackIdentity(model.SlackUserID(ui.User.ID), model.SlackTeamID(ui.Team.ID))
		if err != nil {
			return adapter.SlackAuthenticateRes{}, err
		}
		_, err = a.authrepo.Create(m)
		if err != nil {
			return adapter.SlackAuthenticateRes{}, err
		}
	}
	sa, err := a.authrepo.GetBySlackUserID(model.SlackUserID(ui.User.ID))
	if err != nil {
		return adapter.SlackAuthenticateRes{}, err
	}

	res := adapter.SlackAuthenticateRes{
		UserID: string(sa.UserID),
		Profile: adapter.SlackProfile{
			ID:       ui.User.ID,
			Name:     ui.User.Name,
			ImageURL: ui.User.Image512,
		},
		TeamProfile: adapter.SlackTeamProfile{
			ID:       ui.Team.ID,
			Name:     ui.Team.Name,
			Domain:   ui.Team.Domain,
			ImageURL: ui.Team.Image230,
		},
	}
	return res, nil
}
