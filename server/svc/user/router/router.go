package router

import (
	"github.com/labstack/echo/v4"
	"github.com/onyanko-pon/monorepo/server/svc/post/middleware"
	user "github.com/onyanko-pon/monorepo/server/svc/user/handler"
)

func Router(e *echo.Group) error {
	g := e.Group("/users")
	u, err := user.Init()
	if err != nil {
		return err
	}
	g.GET("/:id", u.Get)
	g.GET("/profile", u.GetProfile, middleware.VerifyMiddleware)
	g.GET("", u.GetAll)
	g.POST("", u.Create, middleware.VerifyMiddleware)
	return nil
}
