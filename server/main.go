package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/mrymam/radio_rec/server/runner"
)

func main() {
	e := echo.New()
	err := runner.NewRouter(e)
	if err != nil {
		msg := fmt.Sprintf("build error: %s", err.Error())
		panic(msg)
	}
}
