package router

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Post(e *echo.Group) error {
	g := e.Group("/posts")

	g.GET("/:id", getPost)
	return nil
}

func getPost(c echo.Context) error {
	id := c.Param("id")
	title := fmt.Sprintf("title %s", id)

	return c.JSON(http.StatusOK, map[string]map[string]string{
		"post": {
			"id":    id,
			"title": title,
		},
	})
}
