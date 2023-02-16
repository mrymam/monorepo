package factory

import (
	"github.com/mrymam/monorepo/server/adapter/svc/authz"
	"github.com/mrymam/monorepo/server/svc/authz/handler"
)

func InitSlackOAuth2() (authz.SlackOAuth2, error) {
	return handler.InitSlackOAuth2()
}

func InitTwitterAuth1() (authz.TwitterAuth1, error) {
	return handler.InitTwitterAuth1()
}
