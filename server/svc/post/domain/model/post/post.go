package post

import "github.com/google/uuid"

type (
	ID    string
	Title string
)

type Post struct {
	ID    ID
	Title Title
}

func Init(title Title) (Post, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return Post{}, err
	}
	return Post{
		ID:    ID(id.String()),
		Title: title,
	}, nil
}
