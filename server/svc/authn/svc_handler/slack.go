package svchandler

import (
	"encoding/json"

	"github.com/onyanko-pon/monorepo/server/svc/authn/domain/model"
	"github.com/onyanko-pon/monorepo/server/svc/authn/infra/repository"
	svcrouter "github.com/onyanko-pon/monorepo/server/svc_router"
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

func (a SlackAuth) Authenticate(arg string) (string, error) {
	var req svcrouter.SlackAuthReq
	err := json.Unmarshal([]byte(arg), &req)
	if err != nil {
		return "", err
	}
	api := slack.New(req.AccessToken)
	ui, err := api.GetUserIdentity()
	if err != nil {
		return "", err
	}
	sa, err := a.authrepo.GetBySlackUserID(model.SlackUserID(ui.User.ID))
	if err != nil {
		return "", err
	}
	res := svcrouter.SlackAuthRes{
		UserID: string(sa.UserID),
	}
	j, err := json.Marshal(res)
	return string(j), err
}
