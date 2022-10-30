package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onyanko-pon/monorepo/server/pkg/http/response"
	"github.com/onyanko-pon/monorepo/server/svc/account/ctx"
	"github.com/onyanko-pon/monorepo/server/svc/account/domain/model/user"
	"github.com/onyanko-pon/monorepo/server/svc/account/infra/repository"
)

type AccountHander struct {
	prepo repository.Profile
}

func Init() (AccountHander, error) {
	prepo, err := repository.InitProfile()
	if err != nil {
		return AccountHander{}, err
	}
	return AccountHander{
		prepo: prepo,
	}, nil
}

type GetProfileRes struct {
	Profile Profile `json:"profile"`
}

func (h AccountHander) Get(c echo.Context) error {
	userid := c.(ctx.AuthContext).GetUserID()
	p, err := h.prepo.GetByUserID(user.ID(userid))
	if err != nil {
		res := response.NewErrorRes(err)
		return c.JSON(http.StatusNotFound, res)
	}
	pst, err := resolveProfile(p)
	if err != nil {
		return err
	}
	res := GetProfileRes{
		Profile: pst,
	}
	return c.JSON(http.StatusOK, res)
}

type GetProfilesRes struct {
	Profiles []Profile `json:"profiles"`
}

func (h AccountHander) GetAll(c echo.Context) error {
	ms, err := h.prepo.GetAll()
	if err != nil {
		res := response.NewErrorRes(err)
		return c.JSON(http.StatusInternalServerError, res)
	}
	ps := []Profile{}
	for _, m := range ms {
		p, err := resolveProfile(m)
		if err != nil {
			res := response.NewErrorRes(err)
			return c.JSON(http.StatusInternalServerError, res)
		}
		ps = append(ps, p)
	}
	res := GetProfilesRes{
		Profiles: ps,
	}
	return c.JSON(http.StatusOK, res)
}

type CreateProfileReq struct {
	Profile Profile `json:"profile"`
}

type CreateUserRes struct {
	Profile Profile `json:"profile"`
}

func (h AccountHander) TwitteSignin(c echo.Context) error {
	var rq CreateProfileReq
	err := c.Bind(&rq)
	if err != nil {
		res := response.NewErrorRes(err)
		return c.JSON(http.StatusNotFound, res)
	}
	p, err := rq.Profile.ToModel()
	if err != nil {
		res := response.NewErrorRes(err)
		return c.JSON(http.StatusNotFound, res)
	}

	p, err = h.prepo.Create(user.ID(userid), p)
	if err != nil {
		res := response.NewErrorRes(err)
		return c.JSON(http.StatusNotFound, res)
	}
	pst, err := resolveProfile(p)
	if err != nil {
		return err
	}
	res := CreateUserRes{
		Profile: pst,
	}
	return c.JSON(http.StatusOK, res)
}
