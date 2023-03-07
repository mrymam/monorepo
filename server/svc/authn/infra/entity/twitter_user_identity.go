package entity

import "github.com/mrymam/radio_rec/server/svc/authn/domain/model"

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

func ToTwitterUserEntity(m model.TwitterUserIdentity) TwitterUserIdentity {
	return TwitterUserIdentity{
		TwitterUserID: string(m.TwitterUserID),
		UserID:        string(m.UserID),
	}
}
