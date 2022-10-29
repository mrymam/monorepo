package router

import (
	svchandler "github.com/onyanko-pon/monorepo/server/svc/authz/svc_handler"
	svcrouter "github.com/onyanko-pon/monorepo/server/svc_router"
)

func SvcRouter() error {
	twauth1, err := svchandler.InitTwitterAuth1()
	if err != nil {
		return err
	}
	if err := svcrouter.AddHandler(svcrouter.TwitterOAuth1FetchAccessToken, twauth1.GetAccessToken); err != nil {
		return err
	}
	slauth2, err := svchandler.InitSlackOAuth2()
	if err != nil {
		return err
	}
	if err := svcrouter.AddHandler(svcrouter.SlackOAuth2FetchAccessToken, slauth2.GetAccessToken); err != nil {
		return err
	}

	return nil
}
