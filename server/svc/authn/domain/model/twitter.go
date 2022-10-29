package model

type TwitterUserName string
type TwitterUserScreenName string
type TwitterUserProfileImageUrl string
type TwitterUser struct {
	ID              TwitterUserID
	Name            TwitterUserName
	ScreenName      TwitterUserScreenName
	ProfileImageUrl TwitterUserProfileImageUrl
}
