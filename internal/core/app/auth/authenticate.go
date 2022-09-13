package auth

import (
	"context"
	"github.com/golang-jwt/jwt"
	ricardoErr "gitlab.com/ricardo-public/errors/pkg/errors"
	tokens "gitlab.com/ricardo-public/jwt-tools/pkg"
	"gitlab.com/ricardo134/auth-service/internal/core/entities"
	authPort "gitlab.com/ricardo134/auth-service/internal/core/ports/auth"
	customRicardoErr "gitlab.com/ricardo134/auth-service/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

// bcryptCost set as 13 gives approximately 480 ms to hash a password with my work computer
const bcryptCost int = 13

type AuthenticateService interface {
	authPort.Authenticate
}

type authenticateService struct {
	repo          authPort.AuthenticationRepository
	notifier      authPort.RegisterNotifier
	accessSecret  []byte
	refreshSecret []byte
	// Expiration in minutes
	//accessTokenExp uint8
	//refreshTokenExp uint16
}

func NewAuthenticateService(repo authPort.AuthenticationRepository, notifier authPort.RegisterNotifier, accessSecret, refreshSecret []byte) AuthenticateService {
	return authenticateService{
		repo:          repo,
		notifier:      notifier,
		accessSecret:  accessSecret,
		refreshSecret: refreshSecret,
	}
}

func (s authenticateService) Login(ctx context.Context, loginRequest entities.LoginRequest) (*entities.SignedTokenPair, error) {
	user, err := s.repo.EmailExists(ctx, loginRequest.Email)
	if err != nil || (*user == entities.User{}) {
		return nil, ricardoErr.New(ricardoErr.ErrUnauthorized, customRicardoErr.ErrCannotFindUserDescription)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		return nil, ricardoErr.New(ricardoErr.ErrUnauthorized, customRicardoErr.ErrCannotFindUserDescription)
	}

	return generate(strconv.Itoa(int(user.ID)), s.accessSecret, s.refreshSecret), nil
}

func (s authenticateService) Save(ctx context.Context, user entities.User) error {
	existingUser, _ := s.repo.EmailExists(ctx, user.Email)
	if existingUser != nil {
		return ricardoErr.New(ricardoErr.ErrForbidden, "user already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcryptCost)
	if err != nil {
		return ricardoErr.New(ricardoErr.ErrNotFound, "could not hash password")
	}
	user.Password = string(hash)

	_, err = s.repo.Save(ctx, user)
	if err == nil {
		_ = s.notifier.Notify(user)
	}

	return err
}

func (s authenticateService) Refresh(ctx context.Context, token string) (*entities.SignedTokenPair, error) {
	pToken, err := tokens.Parse(token, s.refreshSecret)
	if err != nil {
		return nil, ricardoErr.New(ricardoErr.ErrUnauthorized, customRicardoErr.ErrInvalidTokenDescription)
	}
	rClaims := pToken.Claims.(*tokens.RicardoClaims)

	return generate(rClaims.Subject, s.accessSecret, s.refreshSecret), nil
}

func generate(subject string, accessSecret, refreshSecret []byte) *entities.SignedTokenPair {
	accessTokenClaims := jwt.MapClaims{
		"exp":  time.Now().Add(time.Minute * 15).Unix(),
		"iss":  "aut",
		"sub":  subject,
		"role": "user",
	}
	signedAT, _ := tokens.GenerateHS256SignedToken(accessTokenClaims, accessSecret)

	refreshTokenClaims := jwt.MapClaims{
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
		"iss":  "aut",
		"sub":  subject,
		"role": "user",
	}
	signedRT, _ := tokens.GenerateHS256SignedToken(refreshTokenClaims, refreshSecret)

	return &entities.SignedTokenPair{
		Access:  signedAT,
		Refresh: signedRT,
	}
}
