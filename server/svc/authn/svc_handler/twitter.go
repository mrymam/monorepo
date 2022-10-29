package svchandler

type TwitterAuth struct{}

func (a TwitterAuth) Authenticate(accessToken string, accessSecret string) (string, error) {
	return "", nil
}
