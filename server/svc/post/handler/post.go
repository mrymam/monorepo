package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onyanko-pon/monorepo/server/pkg/http/response"
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
		res := response.NewErrorRes(err)
		return c.JSON(http.StatusNotFound, res)
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

type GetPostsRes struct {
	Posts []Post `json:"posts"`
}

func (h PostHander) GetAll(c echo.Context) error {
	ms, err := h.repo.GetAll()
	if err != nil {
		res := response.NewErrorRes(err)
		return c.JSON(http.StatusInternalServerError, res)
	}
	ps := []Post{}
	for _, m := range ms {
		p, err := resolvePost(m)
		if err != nil {
			res := response.NewErrorRes(err)
			return c.JSON(http.StatusInternalServerError, res)
		}
		ps = append(ps, p)
	}
	res := GetPostsRes{
		Posts: ps,
	}
	return c.JSON(http.StatusOK, res)
}

type CreatePostReq struct {
	Post Post `json:"post"`
}

type CreatePostRes struct {
	Post Post `json:"post"`
}

func (h PostHander) Create(c echo.Context) error {
	var rq CreatePostReq
	err := c.Bind(&rq)
	if err != nil {
		res := response.NewErrorRes(err)
		return c.JSON(http.StatusNotFound, res)
	}
	p, err := rq.Post.ToModel()
	if err != nil {
		res := response.NewErrorRes(err)
		return c.JSON(http.StatusNotFound, res)
	}

	p, err = h.repo.Create(p)
	if err != nil {
		res := response.NewErrorRes(err)
		return c.JSON(http.StatusNotFound, res)
	}
	pst, err := resolvePost(p)
	if err != nil {
		return err
	}
	res := CreatePostRes{
		Post: pst,
	}
	return c.JSON(http.StatusOK, res)
}
