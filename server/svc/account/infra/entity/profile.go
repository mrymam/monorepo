package entity

import (
	"github.com/onyanko-pon/monorepo/server/svc/account/domain/model/profile"
)

type Profile struct {
	ID     string `gorm:"id"`
	UserID string `gorm:"user_id"`
	Name   string `gorm:"Name"`
}

func (p Profile) ToModel() profile.Profile {
	return profile.Profile{
		ID:   profile.ID(p.ID),
		Name: profile.Name(p.Name),
	}
}

func ToProfileEntity(m profile.Profile, userID profile.UserID) Profile {
	return Profile{
		ID:     string(m.ID),
		Name:   string(m.Name),
		UserID: string(userID),
	}
}

func (e Profile) TableName() string {
	return "profiles"
}
