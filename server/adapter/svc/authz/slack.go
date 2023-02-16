package authz

type SlackOAuthCode string
type SlackAccessToken string

type SlackOAuth2 interface {
	GetAccessToken(SlackOAuthCode) (SlackAccessToken, error)
}

// func InitSlackOAuth2() (SlackOAuth2, error) {
// 	s, err := svcimpl.InitSlackAuth2Svc()
// 	if err != nil {
// 		return SlackOAuth2{}, nil
// 	}
// 	return SlackOAuth2{
// 		svc: s,
// 	}, nil
// }
