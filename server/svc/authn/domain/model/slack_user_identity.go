package model

import "github.com/google/uuid"

type SlackUserID string
type SlackTeamID string

type SlackUserIdentity struct {
	SlackUserID SlackUserID
	SlackTeamID SlackTeamID
	UserID      UserID
}

func InitTwitterSlackIdentity(suid SlackUserID, stid SlackTeamID) (SlackUserIdentity, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return SlackUserIdentity{}, err
	}
	return SlackUserIdentity{
		SlackUserID: suid,
		SlackTeamID: stid,
		UserID:      UserID(id.String()),
	}, nil
}
