package model

type TwitterImageURL string
type TwitterScreenName string
type TwitterName string

type TwitterUserProfile struct {
	TwitterUserID TwitterUserID
	ScreenName    TwitterScreenName
	Name          TwitterName
	ImageURL      TwitterImageURL
}
