package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/onyanko-pon/monorepo/server/runner"
	"github.com/onyanko-pon/monorepo/server/svc/authn/router"
)

func main() {
	e := echo.New()
	router.SvcRouter()
	err := runner.NewRouter(e)
	if err != nil {
		msg := fmt.Sprintf("build error: %s", err.Error())
		panic(msg)
	}
}
