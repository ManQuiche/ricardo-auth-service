package auth

import (
	"context"
	errors2 "github.com/pkg/errors"
	ricardoErr "gitlab.com/ricardo-public/errors/pkg/errors"
	"gitlab.com/ricardo134/auth-service/internal/core/entities"
	authPort "gitlab.com/ricardo134/auth-service/internal/core/ports/auth"
	"strconv"
)

type ExternalTokenService interface {
	authPort.TokenAuthenticate
}

type externalTokenService struct {
	tokenRepo     authPort.TokenRepository
	authRepo      authPort.AuthenticationRepository
	accessSecret  []byte
	refreshSecret []byte
	notifier      authPort.RegisterNotifier
}

func NewExternalTokenService(
	tRepo authPort.TokenRepository,
	aRepo authPort.AuthenticationRepository,
	notifier authPort.RegisterNotifier,
	accessSecret,
	refreshSecret []byte,
) ExternalTokenService {
	return &externalTokenService{
		tokenRepo:     tRepo,
		authRepo:      aRepo,
		accessSecret:  accessSecret,
		refreshSecret: refreshSecret,
		notifier:      notifier,
	}
}

func (e externalTokenService) Verify(ctx context.Context, token string) (*entities.SignedTokenPair, error) {
	user, err := e.tokenRepo.Verify(ctx, token)
	if err != nil {
		return nil, ricardoErr.New(ricardoErr.ErrUnauthorized, errors2.Wrap(err, "can't find user").Error())
	}

	existingUser, err := e.authRepo.EmailExists(ctx, user.Email)
	if existingUser == nil {
		existingUser, err = e.authRepo.Save(ctx, *user)
		if err == nil {
			return nil, err
		}

		_ = e.notifier.Notify(*existingUser)
	}

	return generate(strconv.Itoa(int(existingUser.ID)), e.accessSecret, e.refreshSecret), nil
}
