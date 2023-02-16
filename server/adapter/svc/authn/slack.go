package authn

type SlackAccessToken string

type SlackAuth interface {
	Authenticate(accessToken SlackAccessToken) (UserID, SlackProfile, SlackTeamProfile, error)
}

type SlackProfile struct {
	ID       string
	Name     string
	ImageURL string
}

type SlackTeamProfile struct {
	ID       string
	Name     string
	Domain   string
	ImageURL string
}
