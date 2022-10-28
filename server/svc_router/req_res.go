package svcrouter

type TokenVerifyReq struct {
	Token string `json:"token"`
}

type TokenVerifyRes struct {
	UserID   string `json:"user_id"`
	Verified bool   `json:"verified"`
}

type TokenEncodeReq struct {
	UserID string `json:"user_id"`
}

type TokenEncodeRes struct {
	Token string `json:"token"`
}
