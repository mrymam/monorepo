package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onyanko-pon/monorepo/server/svc/post/domain/model/post"
	"github.com/onyanko-pon/monorepo/server/svc/post/infra/repository"
)

type PostHander struct {
	repo repository.Post
}

func Init() (PostHander, error) {
	repo, err := repository.InitPost()
	if err != nil {
		return PostHander{}, err
	}
	return PostHander{
		repo: repo,
	}, nil
}

type GetPostRes struct {
	Post Post `json:"post"`
}

func (h PostHander) Get(c echo.Context) error {
	id := c.Param("id")
	p, err := h.repo.Get(post.ID(id))
	if err != nil {
		return err
	}
	pst, err := resolvePost(p)
	if err != nil {
		return err
	}
	res := GetPostRes{
		Post: pst,
	}
	return c.JSON(http.StatusOK, res)
}
