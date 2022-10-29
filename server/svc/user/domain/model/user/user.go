package user

import "github.com/google/uuid"

type (
	ID       string
	Username string
)

type User struct {
	ID       ID
	Username Username
}

func Init(username Username) (User, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return User{}, err
	}
	return User{
		ID:       ID(id.String()),
		Username: username,
	}, nil
}
