package svc

type OAuthToken string
type OAuthSecret string
type OAuthVerifier string

type AccessToken string
type AccessSecret string

type OAuth1Svc interface {
	FetchAccessToken(OAuthToken, OAuthSecret, OAuthVerifier) (AccessToken, AccessSecret, error)
	VerifyAccessToken(AccessToken, AccessSecret) (string, error)
}
