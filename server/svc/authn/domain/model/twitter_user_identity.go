package model

import "github.com/google/uuid"

type TwitterUserID string
type TwitterUserIdentity struct {
	TwitterUserID TwitterUserID
	UserID        UserID
}

func InitTwitterUserIdentity(tuid TwitterUserID) (TwitterUserIdentity, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return TwitterUserIdentity{}, err
	}
	return TwitterUserIdentity{
		TwitterUserID: tuid,
		UserID:        UserID(id.String()),
	}, nil
}
