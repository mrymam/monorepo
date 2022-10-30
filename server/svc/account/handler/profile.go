package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onyanko-pon/monorepo/server/client/authn"
	"github.com/onyanko-pon/monorepo/server/pkg/http/response"
	"github.com/onyanko-pon/monorepo/server/svc/account/ctx"
	"github.com/onyanko-pon/monorepo/server/svc/account/domain/model/profile"
	"github.com/onyanko-pon/monorepo/server/svc/account/domain/model/user"
	"github.com/onyanko-pon/monorepo/server/svc/account/infra/repository"
)

type AccountHander struct {
	prepo   repository.Profile
	twiauth authn.TwitterAuth
}

func Init() (AccountHander, error) {
	prepo, err := repository.InitProfile()
	if err != nil {
		return AccountHander{}, err
	}
	return AccountHander{
		prepo:   prepo,
		twiauth: authn.TwitterAuthImple{},
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

type TwitterSigninReq struct {
	AccessToken  string `json:"access_token"`
	AccessSecret string `json:"access_secret"`
}

type TwitterSigninRes struct {
	Profile Profile `json:"profile"`
}

func (h AccountHander) TwitteSignin(c echo.Context) error {
	var rq TwitterSigninReq
	err := c.Bind(&rq)
	if err != nil {
		res := response.NewErrorRes(err)
		return c.JSON(http.StatusBadRequest, res)
	}
	ars, err := h.twiauth.Authenticate(rq.AccessToken, rq.AccessSecret)
	if err != nil {
		r := response.NewErrorRes(err)
		return c.JSON(http.StatusBadRequest, r)
	}

	var p profile.Profile
	uid := user.ID(ars.UserID)

	ex, err := h.prepo.Exist(uid)
	if err != nil {
		r := response.NewErrorRes(err)
		return c.JSON(http.StatusBadRequest, r)
	}
	if ex {
		p, err = h.prepo.GetByUserID(uid)
		if err != nil {
			r := response.NewErrorRes(err)
			return c.JSON(http.StatusBadRequest, r)
		}
	} else {
		p, err = profile.Init(profile.Name(ars.TwitterProfile.ScreenName))
		if err != nil {
			res := response.NewErrorRes(err)
			return c.JSON(http.StatusInternalServerError, res)
		}
		p, err = h.prepo.Create(uid, p)
		if err != nil {
			res := response.NewErrorRes(err)
			return c.JSON(http.StatusNotFound, res)
		}
	}

	pst, err := resolveProfile(p)
	if err != nil {
		return err
	}
	res := TwitterSigninRes{
		Profile: pst,
	}
	return c.JSON(http.StatusOK, res)
}
