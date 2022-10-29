package repository

import (
	"github.com/onyanko-pon/monorepo/server/svc/authn/domain/model"
	"github.com/onyanko-pon/monorepo/server/svc/authn/infra/entity"
	"github.com/onyanko-pon/monorepo/server/svc/post/di"
	"gorm.io/gorm"
)

func InitSlackAuthRepo() (SlackAuthRepo, error) {
	db, err := di.GetDB()
	if err != nil {
		return SlackAuthRepo{}, err
	}
	return SlackAuthRepo{db: db}, nil
}

type SlackAuthRepo struct {
	db *gorm.DB
}

func (r SlackAuthRepo) GetBySlackUserID(id model.SlackUserID) (model.SlackAuth, error) {
	var e entity.SlackAuth
	r.db.First(&e, "slack_user_id = ?", id)
	return e.ToModel(), nil
}
