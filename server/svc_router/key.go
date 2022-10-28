package svcrouter

type Key string

const (
	UserTokenVerify Key = "user_token_verify"
	UserTokenEncode Key = "user_token_encode"

	TwitterOAuth1FetchAccessToken Key = "twitter_oauth1_fetch_access_token"
	TwitterOAuth1VerifyUser       Key = "twitter_oauth1_verify_user"
)
