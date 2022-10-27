package entity

import "github.com/onyanko-pon/monorepo/server/svc/post/domain/model/post"

type Post struct {
	ID    string `gorm:"id"`
	Title string `gorm:"title"`
}

func (p Post) ToModel() post.Post {
	return post.Post{
		ID:    post.ID(p.ID),
		Title: post.Title(p.Title),
	}
}

func ToPostEntity(m post.Post) Post {
	return Post{
		ID:    string(m.ID),
		Title: string(m.Title),
	}
}

func (p Post) TableName() string {
	return "posts"
}
