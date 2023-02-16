package middleware

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	authn "github.com/mrymam/monorepo/server/adapter/svc/authn"
	authnf "github.com/mrymam/monorepo/server/adapter/svc/authn/factory"
	"github.com/mrymam/monorepo/server/svc/post/ctx"
)

func VerifyMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := extractToken(c.Request().Header["Authorization"][0])
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		authc, err := authnf.InitAuthn()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		areq := authn.AuthnVerifyTokenReq{
			Token: token,
		}
		rs, err := authc.VerifyToken(areq)

		if err != nil || !rs.Varid {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		ac := ctx.InitAuthConctext(c, rs.UserID)
		if err != nil || !rs.Varid {
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return next(ac)
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
