package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mrymam/radio_rec/server/adapter/svc/authn"
	authnf "github.com/mrymam/radio_rec/server/adapter/svc/authn/factory"
	"github.com/mrymam/radio_rec/server/pkg/http/response"
	"github.com/mrymam/radio_rec/server/svc/account/ctx"
	"github.com/mrymam/radio_rec/server/svc/account/domain/model/profile"
	"github.com/mrymam/radio_rec/server/svc/account/domain/model/user"
	"github.com/mrymam/radio_rec/server/svc/account/infra/repository"
)

type AccountHander struct {
	prepo     repository.Profile
	twiauth   authn.TwitterAuth
	tokenauth authn.Authn
}

func Init() (AccountHander, error) {
	prepo, err := repository.InitProfile()
	if err != nil {
		return AccountHander{}, err
	}
	twiauth, err := authnf.InitTwitterAuth()
	tokenauth, err := authnf.InitAuthn()

	return AccountHander{
		prepo:     prepo,
		twiauth:   twiauth,
		tokenauth: tokenauth,
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

	param := authn.TwitterAuthenticateReq{
		AccessToken:  rq.AccessToken,
		AccessSecret: rq.AccessSecret,
	}
	r, err := h.twiauth.Authenticate(param)
	if err != nil {
		r := response.NewErrorRes(err)
		return c.JSON(http.StatusBadRequest, r)
	}

	var p profile.Profile
	uid := user.ID(r.UserID)

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
		p, err = profile.Init(profile.Name(r.Profile.ScreenName))
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
	pld := authn.AuthnEncodeTokenReq{
		UserID: string(uid),
	}
	rt, err := h.tokenauth.EncodeToken(pld)
	if err != nil {
		res := response.NewErrorRes(err)
		return c.JSON(http.StatusNotFound, res)
	}
	c.Response().Header().Set(echo.HeaderAuthorization, fmt.Sprintf("bearer %s", rt.Token))

	pst, err := resolveProfile(p)
	if err != nil {
		return err
	}
	res := TwitterSigninRes{
		Profile: pst,
	}
	return c.JSON(http.StatusOK, res)
}
