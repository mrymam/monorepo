package model

type ImageURL string
type FirstName string
type LastName string
type RealName string
type DisplayName string

type SlackUserProfile struct {
	SlackUserID SlackUserID
	ImageURL    ImageURL
	FirstName   FirstName
	LastName    LastName
	RealName    RealName
	DisplayName DisplayName
}
