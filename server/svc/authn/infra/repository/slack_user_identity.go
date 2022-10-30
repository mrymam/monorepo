package repository

import (
	"github.com/onyanko-pon/monorepo/server/svc/authn/domain/model"
	"github.com/onyanko-pon/monorepo/server/svc/authn/infra/entity"
	"github.com/onyanko-pon/monorepo/server/svc/post/di"
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
