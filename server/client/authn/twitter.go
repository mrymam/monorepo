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

type User struct{}

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
