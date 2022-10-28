package setting

type Twitter struct {
	OAuth1 OAuth1 `yaml:"oauth1"`
}

type OAuth1 struct {
	ConsumerKey    string `yaml:"consumer_key"`
	ConsumerSecret string `yaml:"consumer_secret"`
	CallbackURL    string `yaml:"callback_url"`
}
