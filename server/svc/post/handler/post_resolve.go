package handler

import "github.com/onyanko-pon/monorepo/server/svc/post/domain/model/post"

type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func resolvePost(p post.Post) (Post, error) {
	return Post{
		ID:    string(p.ID),
		Title: string(p.Title),
	}, nil
}

func (p Post) ToModel() (post.Post, error) {
	return post.Init(post.Title(p.Title))
}
