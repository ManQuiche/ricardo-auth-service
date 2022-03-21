package auth

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt"
	tokens "gitlab.com/ricardo-public/jwt-tools/pkg"
	authPort "ricardo/auth-service/internal/core/ports/auth"
	errors2 "ricardo/auth-service/pkg/errors"
)

type AuthorizeService interface {
	authPort.Authorize
}

func NewAuthorizeService(accessSecret, refreshSecret []byte) AuthorizeService {
	return authorizeService{
		accessSecret:  accessSecret,
		refreshSecret: refreshSecret,
	}
}

type authorizeService struct {
	accessSecret  []byte
	refreshSecret []byte
}

func (a authorizeService) AccessAuthorize(ctx context.Context, accessToken string) (bool, error) {
	return a.authorize(ctx, accessToken, a.accessSecret)
}

func (a authorizeService) RefreshAuthorize(ctx context.Context, refreshToken string) (bool, error) {
	return a.authorize(ctx, refreshToken, a.refreshSecret)
}

func (a authorizeService) authorize(ctx context.Context, token string, key []byte) (bool, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &tokens.RicardoClaims{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New(errors2.InvalidToken)
		}
		return key, nil
	})

	if err != nil {
		return false, err
	}

	return parsedToken.Valid, nil
}
