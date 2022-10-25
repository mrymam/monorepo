package repository

import (
	"github.com/onyanko-pon/monorepo/server/svc/post/domain/model/post"
)

type Post interface {
	Get(id post.ID) (post.Post, error)
}

type PostImple struct{}

func InitPost() (Post, error) {
	return PostImple{}, nil
}

func (p PostImple) Get(id post.ID) (post.Post, error) {
	pst, err := post.Init("title")
	if err != nil {
		return post.Post{}, err
	}
	return pst, nil
}
