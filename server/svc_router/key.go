package svcrouter

type Key string

const (
	UserTokenVerify Key = "user_token_verify"
	UserTokenEncode Key = "user_token_encode"

	TwitterOAuth1FetchAccessToken Key = "twitter_oauth1_fetch_access_token"
	SlackOAuth2FetchAccessToken   Key = "slack_oauth2_fetch_access_token"

	TwitterAuth Key = "twitter_auth"
	SlackAuth   Key = "slack_auth"
)
