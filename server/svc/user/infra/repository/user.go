package repository

import (
	"github.com/onyanko-pon/monorepo/server/svc/user/di"
	"github.com/onyanko-pon/monorepo/server/svc/user/domain/model/user"
	"github.com/onyanko-pon/monorepo/server/svc/user/infra/entity"
	"gorm.io/gorm"
)

type User interface {
	Get(id user.ID) (user.User, error)
	GetAll() ([]user.User, error)
	Create(p user.User) (user.User, error)
}

type UserImple struct {
	db *gorm.DB
}

func InitUser() (User, error) {
	db, err := di.GetDB()
	if err != nil {
		return UserImple{}, err
	}
	return UserImple{db}, nil
}

func (r UserImple) Get(id user.ID) (user.User, error) {

	e := entity.User{}
	err := r.db.First(&e, "id = ?", id).Error
	if err != nil {
		return user.User{}, err
	}
	return e.ToModel(), nil
}

func (r UserImple) GetAll() ([]user.User, error) {
	es := []entity.User{}
	err := r.db.Find(&es).Error
	if err != nil {
		return []user.User{}, err
	}
	ms := []user.User{}
	for _, e := range es {
		ms = append(ms, e.ToModel())
	}
	return ms, nil
}

func (r UserImple) Create(p user.User) (user.User, error) {
	e := entity.ToUserEntity(p)
	err := r.db.Create(&e).Error
	if err != nil {
		return user.User{}, err
	}
	return e.ToModel(), nil
}
