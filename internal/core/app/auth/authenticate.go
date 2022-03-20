package auth

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt"
	tokens "gitlab.com/ricardo-public/jwt-tools/pkg"
	"ricardo/auth-service/internal/core/entities"
	authEntities "ricardo/auth-service/internal/core/entities/auth"
	authPort "ricardo/auth-service/internal/core/ports/auth"
	"strconv"
	"time"
)

type AuthenticateService interface {
	authPort.Authenticate
}

type authenticateService struct {
	repo          authPort.AuthenticationRepository
	accessSecret  []byte
	refreshSecret []byte
	// Expiration in minutes
	//accessTokenExp uint8
	//refreshTokenExp uint16
}

func NewAuthenticateService(repo authPort.AuthenticationRepository, accessSecret, refreshSecret []byte) AuthenticateService {
	return authenticateService{
		repo:          repo,
		accessSecret:  accessSecret,
		refreshSecret: refreshSecret,
	}
}

func (s authenticateService) Login(ctx context.Context, loginRequest entities.LoginRequest) (*authEntities.SignedTokenPair, error) {
	user, err := s.repo.Exists(ctx, loginRequest.Email, loginRequest.Password)
	if err != nil || (*user == entities.User{}) {
		return nil, errors.New("cannot find user")
	}

	accessTokenClaims := jwt.MapClaims{
		"exp": time.Now().Add(time.Minute * 15).Unix(),
		"iss": "aut",
		"sub": strconv.Itoa(int(user.ID)),
	}
	signedAT, _ := tokens.GenerateHS256SignedToken(accessTokenClaims, s.accessSecret)

	refreshTokenClaims := jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 72).Unix(),
		"iss": "aut",
		"sub": strconv.Itoa(int(user.ID)),
	}
	signedRT, _ := tokens.GenerateHS256SignedToken(refreshTokenClaims, s.refreshSecret)

	return &authEntities.SignedTokenPair{
		Access:  signedAT,
		Refresh: signedRT,
	}, nil
}

func (s authenticateService) Save(ctx context.Context, user entities.User) error {
	return s.repo.Save(ctx, user)
}
