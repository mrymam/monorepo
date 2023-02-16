package router

import (
	"github.com/labstack/echo/v4"
	post "github.com/mrymam/monorepo/server/svc/post/handler"
	"github.com/mrymam/monorepo/server/svc/post/middleware"
)

func Router(e *echo.Group) error {
	g := e.Group("/posts")
	p, err := post.Init()
	if err != nil {
		return err
	}
	g.GET("/:id", p.Get)
	g.GET("", p.GetAll)
	g.POST("", p.Create, middleware.VerifyMiddleware)
	return nil
}
