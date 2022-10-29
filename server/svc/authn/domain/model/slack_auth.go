package model

type SlackUserID string
type SlackTeamID string

type SlackAuth struct {
	SlackUserID SlackUserID
	SlackTeamID SlackTeamID
	UserID      UserID
}
