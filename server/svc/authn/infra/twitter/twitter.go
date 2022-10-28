package twitter

import "github.com/onyanko-pon/monorepo/server/svc/authn/domain/model"

type User struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	ScreenName      string `json:"screen_name"`
	ProfileImageUrl string `json:"profile_image_url_https"`
}

func (u User) ToModel() model.TwitterUser {
	return model.TwitterUser{
		ID:              u.ID,
		Name:            u.Name,
		ScreenName:      u.ScreenName,
		ProfileImageUrl: u.ProfileImageUrl,
	}
}
