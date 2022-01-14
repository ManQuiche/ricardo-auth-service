package auth

import (
	"auth-service/internal/core/entities"
	authEntities "auth-service/internal/core/entities/auth"
	authPort "auth-service/internal/core/ports/auth"
	"context"
	"github.com/golang-jwt/jwt"
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
	user, err := s.repo.Exists(ctx, loginRequest.Username, loginRequest.Password)
	if err != nil || (*user == entities.User{}) {
		return nil, err
	}

	accessTokenClaims := jwt.MapClaims{
		"exp": time.Now().Add(time.Minute * 15).Unix(),
		"iss": "aut",
		"sub": user.ID,
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	signedAT, _ := accessToken.SignedString(s.accessSecret)

	refreshTokenClaims := jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 72).Unix(),
		"iss": "aut",
		"sub": user.ID,
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	signedRT, _ := refreshToken.SignedString(s.refreshSecret)

	return &authEntities.SignedTokenPair{
		Access:  signedAT,
		Refresh: signedRT,
	}, nil
}

func (s authenticateService) Save(ctx context.Context, user entities.User) error {
	return s.repo.Save(ctx, user)
}
