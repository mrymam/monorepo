package authn

type AuthnVerifyTokenReq struct {
	Token string
}
type AuthnVerifyTokenRes struct {
	UserID string
	Varid  bool
}

type AuthnEncodeTokenReq struct {
	UserID string
}
type AuthnEncodeTokenRes struct {
	Token string
}

type Authn interface {
	VerifyToken(AuthnVerifyTokenReq) (AuthnVerifyTokenRes, error)
	EncodeToken(AuthnEncodeTokenReq) (AuthnEncodeTokenRes, error)
}
