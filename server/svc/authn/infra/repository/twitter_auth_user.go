package repository

import (
	"github.com/onyanko-pon/monorepo/server/svc/authn/domain/model"
	"github.com/onyanko-pon/monorepo/server/svc/authn/infra/entity"
	"gorm.io/gorm"
)

type TwitterAuthUserRepo struct {
	db gorm.DB
}

func (r TwitterAuthUserRepo) GetByTiwtterUserID(id model.TwitterUserID) (model.TwitterAuthUser, error) {
	var e entity.TwitterAuthUser
	r.db.First(&e, "twitter_user_id = ?", id)
	return e.ToModel(), nil
}
