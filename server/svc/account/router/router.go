package router

import (
	"github.com/labstack/echo/v4"
	account "github.com/mrymam/radio_rec/server/svc/account/handler"
	"github.com/mrymam/radio_rec/server/svc/recommend/middleware"
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
