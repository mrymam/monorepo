package authn

import (
	"encoding/json"

	svcrouter "github.com/onyanko-pon/monorepo/server/svc_router"
)

type TwitterProfile struct {
	ID         string
	Name       string
	ScreenName string
	ImageURL   string
}

type TwitterAuthRes struct {
	UserID         string
	TwitterProfile TwitterProfile
}

type TwitterAuth interface {
	Authenticate(accessToken string, accessSecret string) (TwitterAuthRes, error)
}
type TwitterAuthImple struct{}

func (a TwitterAuthImple) Authenticate(accessToken string, accessSecret string) (TwitterAuthRes, error) {
	rq := svcrouter.TwitterAuthReq{
		AccessToken:  accessToken,
		AccessSecret: accessSecret,
	}
	j, err := json.Marshal(rq)
	if err != nil {
		return TwitterAuthRes{}, err
	}
	vrs, err := svcrouter.Handle(svcrouter.TwitterAuth, string(j))
	if err != nil {
		return TwitterAuthRes{}, err
	}
	var rs svcrouter.TwitterAuthRes
	if err = json.Unmarshal([]byte(vrs), &rs); err != nil {
		return TwitterAuthRes{}, err
	}
	return TwitterAuthRes{
		UserID: rs.UserID,
		TwitterProfile: TwitterProfile{
			ID:         rs.Profile.ID,
			Name:       rs.Profile.Name,
			ScreenName: rs.Profile.ScreenName,
			ImageURL:   rs.Profile.ImageURL,
		},
	}, nil
}
