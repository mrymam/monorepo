package entity

type TwitterAuthUser struct {
	TwitterUserID string `gorm:"twitter_user_id"`
	UserID        string `gorm:"user_id"`
}
