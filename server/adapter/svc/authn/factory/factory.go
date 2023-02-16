package factory

import (
	"github.com/onyanko-pon/monorepo/server/adapter/svc/authn"
	"github.com/onyanko-pon/monorepo/server/svc/authn/handler"
)

func InitSlackAuth() (authn.SlackAuth, error) {
	return handler.InitSlackAuth()
}

func InitAuthn() (authn.Authn, error) {
	return handler.InitAuthn()
}

func InitTwitterAuth() (authn.TwitterAuth, error) {
	return handler.InitTwitterAuth()
}
