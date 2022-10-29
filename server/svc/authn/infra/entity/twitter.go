package entity

import "github.com/onyanko-pon/monorepo/server/svc/authn/domain/model"

type TwitterUser struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	ScreenName      string `json:"screen_name"`
	ProfileImageUrl string `json:"profile_image_url_https"`
}

func (e TwitterUser) ToModel() model.TwitterUser {
	return model.TwitterUser{
		ID:              e.ID,
		Name:            e.Name,
		ScreenName:      e.ScreenName,
		ProfileImageUrl: e.ProfileImageUrl,
	}
}
