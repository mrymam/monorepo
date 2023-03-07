package repository

import (
	"github.com/mrymam/radio_rec/server/svc/authn/domain/model"
	"github.com/mrymam/radio_rec/server/svc/authn/infra/entity"
	"github.com/mrymam/radio_rec/server/svc/post/di"
	"gorm.io/gorm"
)

func InitSlackUserIdentityRepo() (SlackUserIdentyRepo, error) {
	db, err := di.GetDB()
	if err != nil {
		return SlackUserIdentyRepo{}, err
	}
	return SlackUserIdentyRepo{db: db}, nil
}

type SlackUserIdentyRepo struct {
	db *gorm.DB
}

func (r SlackUserIdentyRepo) GetBySlackUserID(id model.SlackUserID) (model.SlackUserIdentity, error) {
	var e entity.SlackUserIdentity
	r.db.First(&e, "slack_user_id = ?", id)
	return e.ToModel(), nil
}

func (r SlackUserIdentyRepo) Exist(id model.SlackUserID) (bool, error) {
	var c int64
	var e entity.SlackUserIdentity
	err := r.db.First(&e, "slack_user_id = ?", id).Count(&c).Error
	if err != nil {
		return false, err
	}
	return c > 0, nil
}

func (r SlackUserIdentyRepo) Create(m model.SlackUserIdentity) (model.SlackUserIdentity, error) {
	e := entity.ToSlackUserEntity(m)
	err := r.db.Create(e).Error
	if err != nil {
		return model.SlackUserIdentity{}, err
	}
	return m, nil
}
