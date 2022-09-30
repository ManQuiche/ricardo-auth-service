package auth

import (
	"context"
	"github.com/golang-jwt/jwt"
	ricardoErr "gitlab.com/ricardo-public/errors/pkg/errors"
	tokens "gitlab.com/ricardo-public/jwt-tools/v2/pkg/token"
	authPort "gitlab.com/ricardo134/auth-service/internal/core/ports/auth"
	customRicardoErr "gitlab.com/ricardo134/auth-service/pkg/errors"
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
			return nil, ricardoErr.New(ricardoErr.ErrForbidden, customRicardoErr.ErrInvalidTokenDescription)
		}
		return key, nil
	})

	if err != nil {
		// FIXME: ricardoErr.ErrInternal please ...
		return false, ricardoErr.New("INTERNAL", err.Error())
	}

	return parsedToken.Valid, nil
}
