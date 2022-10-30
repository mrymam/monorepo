package entity

import "github.com/onyanko-pon/monorepo/server/svc/authn/domain/model"

type TwitterUserIdentity struct {
	TwitterUserID string `gorm:"twitter_user_id"`
	UserID        string `gorm:"user_id"`
}

func (e TwitterUserIdentity) ToModel() model.TwitterUserIdentity {
	return model.TwitterUserIdentity{
		TwitterUserID: model.TwitterUserID(e.TwitterUserID),
		UserID:        model.UserID(e.UserID),
	}
}
