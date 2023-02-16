package authn

type UserID string
type Token string
type Varid bool

type Authn interface {
	VerifyToken(Token) (UserID, Varid, error)
	EncodeToken(UserID) (Token, error)
}
