package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type post struct{}

func Init() (post, error) {
	return post{}, nil
}

func (p post) GetPost(c echo.Context) error {
	id := c.Param("id")
	title := fmt.Sprintf("title %s", id)

	return c.JSON(http.StatusOK, map[string]map[string]string{
		"post": {
			"id":    id,
			"title": title,
		},
	})
}
