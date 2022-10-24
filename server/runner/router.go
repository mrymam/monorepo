package runner

import "github.com/labstack/echo/v4"

type router func(e *echo.Echo) error

func newRouter(e *echo.Echo) error {
	routers := []router{}

	for _, r := range routers {
		err := r(e)
		if err != nil {
			return err
		}
	}

	return nil
}
