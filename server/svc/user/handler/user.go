package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onyanko-pon/monorepo/server/pkg/http/response"
	"github.com/onyanko-pon/monorepo/server/svc/user/ctx"
	"github.com/onyanko-pon/monorepo/server/svc/user/domain/model/user"
	"github.com/onyanko-pon/monorepo/server/svc/user/infra/repository"
)

type UserHander struct {
	repo repository.User
}

func Init() (UserHander, error) {
	repo, err := repository.InitUser()
	if err != nil {
		return UserHander{}, err
	}
	return UserHander{
		repo: repo,
	}, nil
}

type GetUserRes struct {
	User User `json:"user"`
}

func (h UserHander) Get(c echo.Context) error {
	id := c.Param("id")
	p, err := h.repo.Get(user.ID(id))
	if err != nil {
		res := response.NewErrorRes(err)
		return c.JSON(http.StatusNotFound, res)
	}
	pst, err := resolveUser(p)
	if err != nil {
		return err
	}
	res := GetUserRes{
		User: pst,
	}
	return c.JSON(http.StatusOK, res)
}

func (h UserHander) GetProfile(c echo.Context) error {
	id := c.(ctx.AuthContext).GetUserID()
	p, err := h.repo.Get(user.ID(id))
	if err != nil {
		res := response.NewErrorRes(err)
		return c.JSON(http.StatusNotFound, res)
	}
	pst, err := resolveUser(p)
	if err != nil {
		return err
	}
	res := GetUserRes{
		User: pst,
	}
	return c.JSON(http.StatusOK, res)
}

type GetUsersRes struct {
	Users []User `json:"users"`
}

func (h UserHander) GetAll(c echo.Context) error {
	ms, err := h.repo.GetAll()
	if err != nil {
		res := response.NewErrorRes(err)
		return c.JSON(http.StatusInternalServerError, res)
	}
	ps := []User{}
	for _, m := range ms {
		p, err := resolveUser(m)
		if err != nil {
			res := response.NewErrorRes(err)
			return c.JSON(http.StatusInternalServerError, res)
		}
		ps = append(ps, p)
	}
	res := GetUsersRes{
		Users: ps,
	}
	return c.JSON(http.StatusOK, res)
}

type CreateUserReq struct {
	User User `json:"user"`
}

type CreateUserRes struct {
	User User `json:"user"`
}

func (h UserHander) Create(c echo.Context) error {
	var rq CreateUserReq
	err := c.Bind(&rq)
	if err != nil {
		res := response.NewErrorRes(err)
		return c.JSON(http.StatusNotFound, res)
	}
	p, err := rq.User.ToModel()
	if err != nil {
		res := response.NewErrorRes(err)
		return c.JSON(http.StatusNotFound, res)
	}

	p, err = h.repo.Create(p)
	if err != nil {
		res := response.NewErrorRes(err)
		return c.JSON(http.StatusNotFound, res)
	}
	pst, err := resolveUser(p)
	if err != nil {
		return err
	}
	res := CreateUserRes{
		User: pst,
	}
	return c.JSON(http.StatusOK, res)
}
