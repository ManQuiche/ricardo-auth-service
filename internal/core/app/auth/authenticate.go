package auth

import (
	"context"
	"github.com/golang-jwt/jwt"
	ricardoErr "gitlab.com/ricardo-public/errors/pkg/errors"
	tokens "gitlab.com/ricardo-public/jwt-tools/pkg"
	"ricardo/auth-service/internal/core/entities"
	authPort "ricardo/auth-service/internal/core/ports/auth"
	customRicardoErr "ricardo/auth-service/pkg/errors"
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

func (s authenticateService) Login(ctx context.Context, loginRequest entities.LoginRequest) (*entities.SignedTokenPair, error) {
	user, err := s.repo.Exists(ctx, loginRequest.Email, loginRequest.Password)
	if err != nil || (*user == entities.User{}) {
		return nil, ricardoErr.New(ricardoErr.ErrUnauthorized, customRicardoErr.ErrCannotFindUserDescription)
	}

	return s.generate(strconv.Itoa(int(user.ID))), nil
}

func (s authenticateService) Save(ctx context.Context, user entities.User) error {
	return s.repo.Save(ctx, user)
}

func (s authenticateService) Refresh(ctx context.Context, token string) (*entities.SignedTokenPair, error) {
	pToken, err := tokens.Parse(token, s.refreshSecret)
	if err != nil {
		return nil, ricardoErr.New(ricardoErr.ErrUnauthorized, customRicardoErr.ErrInvalidTokenDescription)
	}
	rClaims := pToken.Claims.(*tokens.RicardoClaims)

	return s.generate(rClaims.Subject), nil
}

func (s authenticateService) generate(subject string) *entities.SignedTokenPair {
	accessTokenClaims := jwt.MapClaims{
		"exp":  time.Now().Add(time.Minute * 15).Unix(),
		"iss":  "aut",
		"sub":  subject,
		"role": "user",
	}
	signedAT, _ := tokens.GenerateHS256SignedToken(accessTokenClaims, s.accessSecret)

	refreshTokenClaims := jwt.MapClaims{
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
		"iss":  "aut",
		"sub":  subject,
		"role": "user",
	}
	signedRT, _ := tokens.GenerateHS256SignedToken(refreshTokenClaims, s.refreshSecret)

	return &entities.SignedTokenPair{
		Access:  signedAT,
		Refresh: signedRT,
	}
}
