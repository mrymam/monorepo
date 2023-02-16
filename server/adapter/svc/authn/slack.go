package authn

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

type SlackAuthenticateReq struct {
	AccessToken string
}

type SlackAuthenticateRes struct {
	UserID      string
	Profile     SlackProfile
	TeamProfile SlackTeamProfile
}

type SlackAuth interface {
	Authenticate(SlackAuthenticateReq) (SlackAuthenticateRes, error)
}
