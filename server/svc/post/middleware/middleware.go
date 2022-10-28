package middleware

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/onyanko-pon/monorepo/server/client/authn"
)

func VerifyMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := extractToken(c.Request().Header["Authorization"][0])
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		clt := authn.TokenImple{}
		rs, err := clt.VerifyToken(token)

		if err != nil || !rs.Verified {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		return next(c)
	}
}

func extractToken(token string) (string, error) {
	if token == "" {
		return "", nil
	}

	if token[:7] != "Bearer " {
		return "", fmt.Errorf("invalid Bearer token")
	}
	return token[7:], nil
}
