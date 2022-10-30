package model

type SlackUserID string
type SlackTeamID string

type SlackUserIdentity struct {
	SlackUserID SlackUserID
	SlackTeamID SlackTeamID
	UserID      UserID
}
