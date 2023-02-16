package handler

import (
	adapter "github.com/onyanko-pon/monorepo/server/adapter/svc/authn"
	"github.com/onyanko-pon/monorepo/server/svc/authn/domain/model"
	"github.com/onyanko-pon/monorepo/server/svc/authn/infra/repository"
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

func (a SlackAuth) Authenticate(accessToken adapter.SlackAccessToken) (adapter.UserID, adapter.SlackProfile, adapter.SlackTeamProfile, error) {
	api := slack.New(string(accessToken))
	ui, err := api.GetUserIdentity()
	if err != nil {
		return "", adapter.SlackProfile{}, adapter.SlackTeamProfile{}, err
	}
	ex, err := a.authrepo.Exist(model.SlackUserID(ui.User.ID))
	if err != nil {
		return "", adapter.SlackProfile{}, adapter.SlackTeamProfile{}, err
	}
	if !ex {
		m, err := model.InitTwitterSlackIdentity(model.SlackUserID(ui.User.ID), model.SlackTeamID(ui.Team.ID))
		if err != nil {
			return "", adapter.SlackProfile{}, adapter.SlackTeamProfile{}, err
		}
		_, err = a.authrepo.Create(m)
		if err != nil {
			return "", adapter.SlackProfile{}, adapter.SlackTeamProfile{}, err
		}
	}
	sa, err := a.authrepo.GetBySlackUserID(model.SlackUserID(ui.User.ID))
	if err != nil {
		return "", adapter.SlackProfile{}, adapter.SlackTeamProfile{}, err
	}

	return adapter.UserID(sa.UserID), adapter.SlackProfile{
			ID:       ui.User.ID,
			Name:     ui.User.Name,
			ImageURL: ui.User.Image512,
		},
		adapter.SlackTeamProfile{
			ID:       ui.Team.ID,
			Name:     ui.Team.Name,
			Domain:   ui.Team.Domain,
			ImageURL: ui.Team.Image230,
		},
		nil
}
