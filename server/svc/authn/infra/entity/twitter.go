package entity

import "github.com/mrymam/radio_rec/server/svc/authn/domain/model"

type TwitterUser struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	ScreenName      string `json:"screen_name"`
	ProfileImageUrl string `json:"profile_image_url_https"`
}

func (e TwitterUser) ToModel() model.TwitterUser {
	return model.TwitterUser{
		ID:              model.TwitterUserID(e.ID),
		Name:            model.TwitterUserName(e.Name),
		ScreenName:      model.TwitterUserScreenName(e.ScreenName),
		ProfileImageUrl: model.TwitterUserProfileImageUrl(e.ProfileImageUrl),
	}
}
