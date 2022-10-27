package svcrouter

type UserVerifyReq struct {
	Token string `json:"token"`
}

type UserVerifyRes struct {
	UserID   string `json:"user_id"`
	Verified bool   `json:"verified"`
}
