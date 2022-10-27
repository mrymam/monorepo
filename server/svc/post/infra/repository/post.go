package repository

import (
	"github.com/onyanko-pon/monorepo/server/svc/post/di"
	"github.com/onyanko-pon/monorepo/server/svc/post/domain/model/post"
	"github.com/onyanko-pon/monorepo/server/svc/post/infra/entity"
	"gorm.io/gorm"
)

type Post interface {
	Get(id post.ID) (post.Post, error)
}

type PostImple struct {
	db *gorm.DB
}

func InitPost() (Post, error) {
	db, err := di.GetDB()
	if err != nil {
		return PostImple{}, err
	}
	return PostImple{db}, nil
}

func (r PostImple) Get(id post.ID) (post.Post, error) {

	e := entity.Post{}
	err := r.db.First(&e, "id = ?", id).Error
	if err != nil {
		return post.Post{}, err
	}
	return e.ToModel(), nil
}

func (r PostImple) Create(p post.Post) (post.Post, error) {
	e := entity.ToPostEntity(p)
	err := r.db.Create(&e).Error
	if err != nil {
		return post.Post{}, err
	}
	return e.ToModel(), nil
}
