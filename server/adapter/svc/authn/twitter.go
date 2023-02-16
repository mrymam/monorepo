package authn

type TwitterProfile struct {
	ID         string `json:"id"`
	ScreenName string `json:"screen_name"`
	Name       string `json:"name"`
	ImageURL   string `json:"image_url"`
}

type TwitterAuthenticateReq struct {
	AccessToken  string
	AccessSecret string
}

type TwitterAuthenticateRes struct {
	UserID  string
	Profile TwitterProfile
}

type TwitterAuth interface {
	Authenticate(TwitterAuthenticateReq) (TwitterAuthenticateRes, error)
}
