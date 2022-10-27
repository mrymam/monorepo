package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onyanko-pon/monorepo/server/client/authn"
)

func VerifyMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := "token"
		clt := authn.AuthnImpl{}
		rs, err := clt.Verify(token)

		if err != nil || !rs.Verified {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		return next(c)
	}
}
