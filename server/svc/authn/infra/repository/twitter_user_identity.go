package repository

import (
	"github.com/onyanko-pon/monorepo/server/svc/authn/domain/model"
	"github.com/onyanko-pon/monorepo/server/svc/authn/infra/entity"
	"github.com/onyanko-pon/monorepo/server/svc/post/di"
	"gorm.io/gorm"
)

func InitTwitterUserIdentityRepo() (TwitterUserIdentityRepo, error) {
	db, err := di.GetDB()
	if err != nil {
		return TwitterUserIdentityRepo{}, err
	}
	return TwitterUserIdentityRepo{db: db}, nil
}

type TwitterUserIdentityRepo struct {
	db *gorm.DB
}

func (r TwitterUserIdentityRepo) GetByTiwtterUserID(id model.TwitterUserID) (model.TwitterUserIdentity, error) {
	var e entity.TwitterUserIdentity
	r.db.First(&e, "twitter_user_id = ?", id)
	return e.ToModel(), nil
}
