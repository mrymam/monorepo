package authn

type TwitterProfile struct {
	ID         string `json:"id"`
	ScreenName string `json:"screen_name"`
	Name       string `json:"name"`
	ImageURL   string `json:"image_url"`
}

type TwitterAccessToken string
type TwitterAccessSecret string

type TwitterAuth interface {
	Authenticate(TwitterAccessToken, TwitterAccessSecret) (UserID, TwitterProfile, error)
}
