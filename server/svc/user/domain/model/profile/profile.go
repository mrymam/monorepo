package profile

import "github.com/google/uuid"

type (
	ID     string
	UserID string
	Name   string
)

type Profile struct {
	ID   ID
	Name Name
}

func Init(name Name) (Profile, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return Profile{}, err
	}
	return Profile{
		ID:   ID(id.String()),
		Name: name,
	}, nil
}
