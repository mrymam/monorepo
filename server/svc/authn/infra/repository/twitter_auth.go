package repository

import (
	"github.com/onyanko-pon/monorepo/server/svc/authn/domain/model"
	"github.com/onyanko-pon/monorepo/server/svc/authn/infra/entity"
	"github.com/onyanko-pon/monorepo/server/svc/post/di"
	"gorm.io/gorm"
)

func InitTwitterAuthRepo() (TwitterAuthRepo, error) {
	db, err := di.GetDB()
	if err != nil {
		return TwitterAuthRepo{}, err
	}
	return TwitterAuthRepo{db: db}, nil
}

type TwitterAuthRepo struct {
	db *gorm.DB
}

func (r TwitterAuthRepo) GetByTiwtterUserID(id model.TwitterUserID) (model.TwitterAuth, error) {
	var e entity.TwitterAuth
	r.db.First(&e, "twitter_user_id = ?", id)
	return e.ToModel(), nil
}
