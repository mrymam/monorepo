package router

import (
	"github.com/labstack/echo/v4"
	account "github.com/onyanko-pon/monorepo/server/svc/account/handler"
	"github.com/onyanko-pon/monorepo/server/svc/post/middleware"
)

func Router(e *echo.Group) error {
	g := e.Group("/accounts")
	u, err := account.Init()
	if err != nil {
		return err
	}
	g.GET("/profile", u.Get, middleware.VerifyMiddleware)
	g.GET("", u.GetAll)
	g.POST("/twitter_signin", u.TwitteSignin)
	return nil
}
