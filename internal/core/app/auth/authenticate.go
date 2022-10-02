package auth

import (
	"context"
	errorsext "gitlab.com/ricardo-public/errors/pkg/errors"
	tokens "gitlab.com/ricardo-public/jwt-tools/v2/pkg/token"
	"gitlab.com/ricardo134/auth-service/internal/core/entities"
	authPort "gitlab.com/ricardo134/auth-service/internal/core/ports/auth"
	"gitlab.com/ricardo134/auth-service/internal/core/ports/user"
	customRicardoErr "gitlab.com/ricardo134/auth-service/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

const (
	// bcryptCost set as 13 gives approximately 480 ms to hash a password with my work computer
	bcryptCost int = 13
)

type AuthenticateService interface {
	authPort.Authenticate
}

type authenticateService struct {
	repo          authPort.AuthenticationRepository
	notifier      user.EventsNotifier
	accessSecret  []byte
	refreshSecret []byte
	// Expiration in minutes
	//accessTokenExp uint8
	//refreshTokenExp uint16
}

func NewAuthenticateService(
	repo authPort.AuthenticationRepository,
	notifier user.EventsNotifier,
	accessSecret, refreshSecret []byte,
) AuthenticateService {
	return authenticateService{
		repo:          repo,
		notifier:      notifier,
		accessSecret:  accessSecret,
		refreshSecret: refreshSecret,
	}
}

func (s authenticateService) Login(ctx context.Context, loginRequest entities.LoginRequest) (*tokens.SignedTokens, error) {
	user, err := s.repo.EmailExists(ctx, loginRequest.Email)
	if err != nil || (*user == entities.User{}) {
		return nil, errorsext.New(errorsext.ErrUnauthorized, customRicardoErr.ErrCannotFindUserDescription)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		return nil, errorsext.New(errorsext.ErrUnauthorized, customRicardoErr.ErrCannotFindUserDescription)
	}

	return generate(strconv.Itoa(int(user.ID)), user.Role, s.accessSecret, s.refreshSecret), nil
}

func (s authenticateService) Save(ctx context.Context, user entities.User) error {
	existingUser, _ := s.repo.EmailExists(ctx, user.Email)
	if existingUser != nil {
		return errorsext.New(errorsext.ErrForbidden, "user already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcryptCost)
	if err != nil {
		return errorsext.New(errorsext.ErrNotFound, "could not hash password")
	}
	user.Password = string(hash)

	createdUser, err := s.repo.Save(ctx, user)
	if err == nil {
		_ = s.notifier.Created(*createdUser)
	}

	return err
}

func (s authenticateService) Refresh(ctx context.Context, token string) (*tokens.SignedTokens, error) {
	pToken, err := tokens.Parse(token, s.refreshSecret)
	if err != nil {
		return nil, errorsext.New(errorsext.ErrUnauthorized, customRicardoErr.ErrInvalidTokenDescription)
	}
	rClaims := pToken.Claims.(*tokens.RicardoClaims)

	return generate(rClaims.Subject, tokens.Role(rClaims.Role), s.accessSecret, s.refreshSecret), nil
}

func generate(subject string, role tokens.Role, acSec, reSec []byte) *tokens.SignedTokens {
	acClaims := tokens.NewRicardoClaims(subject, "aut", role, time.Now().Add(time.Minute*15))
	signedAT, _ := tokens.GenerateHS256SignedToken(acClaims, acSec)

	reClaims := tokens.NewRicardoClaims(subject, "aut", role, time.Now().Add(time.Minute*72))
	signedRT, _ := tokens.GenerateHS256SignedToken(reClaims, reSec)

	return &tokens.SignedTokens{
		Access:  signedAT,
		Refresh: signedRT,
	}
}
