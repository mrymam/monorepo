package entity

import "github.com/onyanko-pon/monorepo/server/svc/authn/domain/model"

type SlackUserIdentity struct {
	SlackUserID string `gorm:"slack_user_id"`
	SlackTeamID string `gorm:"slack_team_id"`
	UserID      string `gorm:"user_id"`
}

func (e SlackUserIdentity) ToModel() model.SlackUserIdentity {
	return model.SlackUserIdentity{
		SlackUserID: model.SlackUserID(e.SlackUserID),
		SlackTeamID: model.SlackTeamID(e.SlackTeamID),
		UserID:      model.UserID(e.UserID),
	}
}
