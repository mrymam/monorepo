package ctx

import (
	"github.com/labstack/echo/v4"
)

type AuthContext struct {
	echo.Context
	userID string
}

func InitAuthConctext(c echo.Context, userID string) AuthContext {
	return AuthContext{c, userID}
}

func (c AuthContext) GetUserID() string {
	return c.userID
}
