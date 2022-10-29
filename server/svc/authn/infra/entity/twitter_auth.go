package entity

import "github.com/onyanko-pon/monorepo/server/svc/authn/domain/model"

type TwitterAuth struct {
	TwitterUserID string `gorm:"twitter_user_id"`
	UserID        string `gorm:"user_id"`
}

func (e TwitterAuth) ToModel() model.TwitterAuth {
	return model.TwitterAuth{
		TwitterUserID: model.TwitterUserID(e.TwitterUserID),
		UserID:        model.UserID(e.UserID),
	}
}
