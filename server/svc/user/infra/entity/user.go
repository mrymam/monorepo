package entity

import (
	"github.com/onyanko-pon/monorepo/server/svc/user/domain/model/user"
)

type User struct {
	ID       string `gorm:"id"`
	Username string `gorm:"username"`
}

func (u User) ToModel() user.User {
	return user.User{
		ID:       user.ID(u.ID),
		Username: user.Username(u.Username),
	}
}

func ToUserEntity(m user.User) User {
	return User{
		ID:       string(m.ID),
		Username: string(m.Username),
	}
}

func (e User) TableName() string {
	return "users"
}
