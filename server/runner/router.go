package runner

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type router func(e *echo.Echo) error

func NewRouter(e *echo.Echo) error {
	routers := []router{}

	for _, r := range routers {
		err := r(e)
		if err != nil {
			return err
		}
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Logger.Fatal(e.Start(":1323"))

	return nil
}
