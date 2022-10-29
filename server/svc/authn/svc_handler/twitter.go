package svchandler

import (
	"encoding/json"

	"github.com/dghubble/oauth1"
	"github.com/onyanko-pon/monorepo/server/svc/authn/infra/repository"
	svcrouter "github.com/onyanko-pon/monorepo/server/svc_router"
)

type TwitterAuth struct {
	twrepo   repository.TwitterUserRepo
	authrepo repository.TwitterAuthRepo
}

func InitTwitterAuth() (TwitterAuth, error) {
	twrepo, err := repository.InitTwitterUserRepo()
	if err != nil {
		return TwitterAuth{}, err
	}
	authrepo, err := repository.InitTwitterAuthRepo()
	if err != nil {
		return TwitterAuth{}, err
	}
	return TwitterAuth{
		twrepo:   twrepo,
		authrepo: authrepo,
	}, nil
}

func (a TwitterAuth) Authenticate(arg string) (string, error) {
	var req svcrouter.TwitterAuthReq
	err := json.Unmarshal([]byte(arg), &req)
	if err != nil {
		return "", err
	}
	token := oauth1.NewToken(req.AccessToken, req.AccessSecret)
	u, err := a.twrepo.Get(token)
	if err != nil {
		return "", err
	}
	at, err := a.authrepo.GetByTiwtterUserID(u.ID)
	if err != nil {
		return "", err
	}

	rs := svcrouter.TwitterAuthRes{
		UserID: string(at.TwitterUserID),
	}
	j, err := json.Marshal(rs)
	return string(j), err
}
