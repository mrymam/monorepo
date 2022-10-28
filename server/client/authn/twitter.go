package authn

import (
	"encoding/json"

	svcrouter "github.com/onyanko-pon/monorepo/server/svc_router"
)

type TwitterOAuth0 interface {
	FetchAccessToken(OAuthToken) (AccessToken, error)
	VerifyAccessToken(AccessToken) (User, error)
}

type OAuthToken struct {
	Token    string
	Secret   string
	Verifier string
}

type AccessToken struct {
	Token  string
	Secret string
}

type TwitterOAuth0Impl struct{}

type User struct {
	ID              string
	Name            string
	ProfileImageURL string
	ScreenName      string
}

func (o TwitterOAuth0Impl) FetchAccessToken(t OAuthToken) (AccessToken, error) {
	q := svcrouter.TwitterOAuth1GetAccessTokenReq{
		OAuthToken:    t.Token,
		OAuthSecret:   t.Secret,
		OAuthVerifier: t.Verifier,
	}

	j, err := json.Marshal(q)
	if err != nil {
		return AccessToken{}, err
	}
	r, err := svcrouter.Handle(svcrouter.TwitterOAuth1FetchAccessToken, string(j))
	if err != nil {
		return AccessToken{}, err
	}

	var rs svcrouter.TwitterOAuth1GetAccessTokenRes
	err = json.Unmarshal([]byte(r), &rs)
	if err != nil {
		return AccessToken{}, err
	}
	return AccessToken{
		Token:  rs.AccessToken,
		Secret: rs.AccessSecret,
	}, nil

}

func (o TwitterOAuth0Impl) VerifyAccessToken(t AccessToken) (User, error) {
	q := svcrouter.TwitterOAuth1VerifyAccessTokenReq{
		AccessToken:  t.Token,
		AccessSecret: t.Secret,
	}
	j, err := json.Marshal(q)
	if err != nil {
		return User{}, err
	}
	r, err := svcrouter.Handle(svcrouter.TwitterOAuth1VerifyAccessToken, string(j))
	if err != nil {
		return User{}, err
	}

	var rs svcrouter.TwitterOAuth1VerifyAccessTokenRes
	err = json.Unmarshal([]byte(r), &rs)
	if err != nil {
		return User{}, err
	}
	return User{
		ID:              rs.User.ID,
		Name:            rs.User.Name,
		ScreenName:      rs.User.ScreenName,
		ProfileImageURL: rs.User.ProfileImageUrl,
	}, nil

}
