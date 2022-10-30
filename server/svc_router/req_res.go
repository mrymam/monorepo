package svcrouter

type TokenVerifyReq struct {
	Token string `json:"token"`
}

type TokenVerifyRes struct {
	UserID   string `json:"user_id"`
	Verified bool   `json:"verified"`
}

type TokenEncodeReq struct {
	UserID string `json:"user_id"`
}

type TokenEncodeRes struct {
	Token string `json:"token"`
}

type TwitterOAuth1GetAccessTokenReq struct {
	OAuthToken    string `json:"oauth_token"`
	OAuthSecret   string `json:"oauth_secret"`
	OAuthVerifier string `json:"oauth_verifier"`
}

type TwitterOAuth1GetAccessTokenRes struct {
	AccessToken  string `json:"access_token"`
	AccessSecret string `json:"access_secret"`
}

type SlackOAuth2GetAccessTokenReq struct {
	OAuthCode string `json:"oauth_code"`
}

type SlackOAuth2GetAccessTokenRes struct {
	AccessToken string `json:"access_token"`
}

type TwitterAuthReq struct {
	AccessToken  string `json:"access_token"`
	AccessSecret string `json:"access_secret"`
}

type TwitterAuthRes struct {
	UserID  string         `json:"user_id"`
	Profile TwitterProfile `json:"profile"`
}

type TwitterProfile struct {
	ID         string `json:"id"`
	ScreenName string `json:"screen_name"`
	Name       string `json:"name"`
	ImageURL   string `json:"image_url"`
}

type SlackAuthReq struct {
	AccessToken string `json:"access_token"`
}

type SlackAuthRes struct {
	UserID  string           `json:"user_id"`
	Profile SlackProfile     `json:"profile"`
	Team    SlackTeamProfile `json:"team"`
}

type SlackProfile struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type SlackTeamProfile struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Domain   string `json:"domain"`
	ImageURL string `json:"image_url"`
}
