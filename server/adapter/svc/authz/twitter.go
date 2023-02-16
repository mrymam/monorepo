package authz

type TwitterOAuthToken string
type TwitterOAuthSecret string
type TwitterOAuthVerifier string

type TwitterAccessToken string
type TwitterAccessSecret string

type TwitterAuth1 interface {
	GetAccessToken(TwitterOAuthToken, TwitterOAuthSecret, TwitterOAuthVerifier) (TwitterAccessToken, TwitterAccessSecret, error)
}
