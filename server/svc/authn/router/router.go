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

	if err := svcrouter.AddHandler(svcrouter.UserVerify, authn.Verify); err != nil {
		return err
	}
	return nil
}
