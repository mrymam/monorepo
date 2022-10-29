package handler

import (
	"github.com/onyanko-pon/monorepo/server/svc/user/domain/model/user"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func resolveUser(p user.User) (User, error) {
	return User{
		ID:       string(p.ID),
		Username: string(p.Username),
	}, nil
}

func (u User) ToModel() (user.User, error) {
	return user.Init(user.Username(u.Username))
}
