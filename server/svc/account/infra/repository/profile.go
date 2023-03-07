package repository

import (
	"github.com/mrymam/radio_rec/server/svc/account/di"
	"github.com/mrymam/radio_rec/server/svc/account/domain/model/profile"
	"github.com/mrymam/radio_rec/server/svc/account/domain/model/user"
	"github.com/mrymam/radio_rec/server/svc/account/infra/entity"
	"gorm.io/gorm"
)

type Profile interface {
	GetByUserID(id user.ID) (profile.Profile, error)
	GetAll() ([]profile.Profile, error)
	Create(user.ID, profile.Profile) (profile.Profile, error)
	Exist(uid user.ID) (bool, error)
}

type ProfileImple struct {
	db *gorm.DB
}

func InitProfile() (Profile, error) {
	db, err := di.GetDB()
	if err != nil {
		return ProfileImple{}, err
	}
	return ProfileImple{db}, nil
}

func (r ProfileImple) GetByUserID(id user.ID) (profile.Profile, error) {

	e := entity.Profile{}
	err := r.db.First(&e, "user_id = ?", id).Error
	if err != nil {
		return profile.Profile{}, err
	}
	return e.ToModel(), nil
}

func (r ProfileImple) GetAll() ([]profile.Profile, error) {
	es := []entity.Profile{}
	err := r.db.Find(&es).Error
	if err != nil {
		return []profile.Profile{}, err
	}
	ms := []profile.Profile{}
	for _, e := range es {
		ms = append(ms, e.ToModel())
	}
	return ms, nil
}

func (r ProfileImple) Create(uid user.ID, p profile.Profile) (profile.Profile, error) {
	e := entity.ToProfileEntity(p, profile.UserID(uid))
	err := r.db.Create(&e).Error
	if err != nil {
		return profile.Profile{}, err
	}
	return e.ToModel(), nil
}

func (r ProfileImple) Exist(uid user.ID) (bool, error) {
	var c int64
	var e entity.Profile
	err := r.db.First(&e, "user_id = ?", uid).Count(&c).Error
	if err != nil {
		return false, err
	}
	return c > 0, nil
}
