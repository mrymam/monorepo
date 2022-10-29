package entity

import "github.com/onyanko-pon/monorepo/server/svc/authn/domain/model"

type TwitterAuthUser struct {
	TwitterUserID string `gorm:"twitter_user_id"`
	UserID        string `gorm:"user_id"`
}

func (e TwitterAuthUser) ToModel() model.TwitterAuthUser {
	return model.TwitterAuthUser{
		TwitterUserID: model.TwitterUserID(e.TwitterUserID),
		UserID:        model.UserID(e.UserID),
	}
}
