package entity

import "github.com/onyanko-pon/monorepo/server/svc/authn/domain/model"

type SlackAuth struct {
	SlackUserID string `gorm:"slack_user_id"`
	SlackTeamID string `gorm:"slack_team_id"`
	UserID      string `gorm:"user_id"`
}

func (e SlackAuth) ToModel() model.SlackAuth {
	return model.SlackAuth{
		SlackUserID: model.SlackUserID(e.SlackUserID),
		SlackTeamID: model.SlackTeamID(e.SlackTeamID),
		UserID:      model.UserID(e.UserID),
	}
}
