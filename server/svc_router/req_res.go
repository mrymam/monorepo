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

type TwitterOAuth1VerifyUserReq struct {
	AccessToken  string `json:"access_token"`
	AccessSecret string `json:"access_secret"`
}

type TwitterOAuth1VerifyUserRes struct {
	User TwitterUser `json:"user"`
}

type TwitterUser struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	ScreenName      string `json:"screen_name"`
	ProfileImageUrl string `json:"profile_image_url_https"`
}
