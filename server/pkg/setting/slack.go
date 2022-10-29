package setting

type Slack struct {
	OAuth2 OAuth2 `yaml:"oauth2"`
}

type OAuth2 struct {
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
	RedirectURL  string `yaml:"redirect_url"`
}
