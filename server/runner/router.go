package runner

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	rtr "github.com/onyanko-pon/monorepo/server/svc/router"
)

type router func(e *echo.Group) error

func NewRouter(e *echo.Echo) error {
	routers := []router{
		rtr.Post,
	}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	api := e.Group("/api")

	for _, r := range routers {
		err := r(api)
		if err != nil {
			return err
		}
	}

	e.Logger.Fatal(e.Start(":1323"))

	return nil
}
