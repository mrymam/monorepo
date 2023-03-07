package repository

import (
	"github.com/mrymam/radio_rec/server/svc/authn/domain/model"
	"github.com/mrymam/radio_rec/server/svc/authn/infra/entity"
	"github.com/mrymam/radio_rec/server/svc/post/di"
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
	err := r.db.First(&e, "twitter_user_id = ?", id).Error
	return e.ToModel(), err
}

func (r TwitterUserIdentityRepo) Exist(id model.TwitterUserID) (bool, error) {
	var c int64
	var e entity.TwitterUserIdentity
	err := r.db.First(&e, "twitter_user_id = ?", id).Count(&c).Error
	if err != nil {
		return false, err
	}
	return c > 0, nil
}

func (r TwitterUserIdentityRepo) Create(m model.TwitterUserIdentity) (model.TwitterUserIdentity, error) {
	e := entity.ToTwitterUserEntity(m)
	err := r.db.Create(e).Error
	if err != nil {
		return model.TwitterUserIdentity{}, err
	}
	return m, nil
}
