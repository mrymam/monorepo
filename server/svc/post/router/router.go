package router

import (
	"github.com/labstack/echo/v4"
	post "github.com/onyanko-pon/monorepo/server/svc/post/handler"
)

func Router(e *echo.Group) error {
	g := e.Group("/posts")
	p, err := post.Init()
	if err != nil {
		return err
	}
	g.GET("/:id", p.Get)
	return nil
}
