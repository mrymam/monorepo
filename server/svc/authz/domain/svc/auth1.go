package svc

type OAuthToken string
type OAuthSecret string
type OAuthVerifier string

type AccessToken string
type AccessSecret string

type OAuth1Svc interface {
	FetchAccessToken(OAuthToken, OAuthSecret, OAuthVerifier) (AccessToken, AccessSecret, error)
}

type OAuth2Svc interface {
	GetAccessToken(authnCode string) (string, error)
}
