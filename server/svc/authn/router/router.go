package router

import (
	svchandler "github.com/onyanko-pon/monorepo/server/svc/authn/svc_handler"
	svcrouter "github.com/onyanko-pon/monorepo/server/svc_router"
)

func SvcRouter() error {
	authn, err := svchandler.InitAuthn()
	if err != nil {
		return err
	}

	if err := svcrouter.AddHandler(svcrouter.UserTokenVerify, authn.VerifyToken); err != nil {
		return err
	}
	if err := svcrouter.AddHandler(svcrouter.UserTokenEncode, authn.EncodeToken); err != nil {
		return err
	}

	slackauth, err := svchandler.InitSlackAuth()
	if err != nil {
		return err
	}
	if err := svcrouter.AddHandler(svcrouter.SlackAuth, slackauth.Authenticate); err != nil {
		return err
	}
	twitterauth, err := svchandler.InitTwitterAuth()
	if err != nil {
		return err
	}
	if err := svcrouter.AddHandler(svcrouter.TwitterAuth, twitterauth.Authenticate); err != nil {
		return err
	}
	return nil
}
