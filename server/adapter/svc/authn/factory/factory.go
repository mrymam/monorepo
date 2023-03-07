package factory

import (
	"github.com/mrymam/radio_rec/server/adapter/svc/authn"
	"github.com/mrymam/radio_rec/server/svc/authn/handler"
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
