package svc

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/onyanko-pon/monorepo/server/svc/authn/domain/model"
	"github.com/onyanko-pon/monorepo/server/svc/authn/domain/svc"
	"github.com/pkg/errors"
)

type JwtSvcImple struct {
	secretKey   string
	validPeriod time.Duration
}

func (v JwtSvcImple) Parse(tokenStr string) (svc.Payload, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(v.secretKey), nil
	})
	if err != nil {
		return svc.Payload{}, fmt.Errorf("invalid jwt token: %s", err.Error())
	}

	c, ok := token.Claims.(jwt.StandardClaims)
	if !ok || token.Valid {
		return svc.Payload{}, fmt.Errorf("invalid jwt token")
	}
	return v.parseClaims(c), nil
}

func (v JwtSvcImple) parseClaims(c jwt.StandardClaims) svc.Payload {
	return svc.Payload{
		UserID: model.UserID(c.Subject),
	}
}

func (v JwtSvcImple) Encode(p svc.Payload) (string, error) {
	now := time.Now()
	claims := jwt.StandardClaims{
		Subject:   string(p.UserID),
		IssuedAt:  now.Unix(),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenStr, err := token.SignedString(v.secretKey)
	if err != nil {
		return "", errors.Wrap(err, "Failed to sign token")
	}
	return tokenStr, nil
}
