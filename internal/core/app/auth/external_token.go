package auth

import (
	"context"
	errors2 "github.com/pkg/errors"
	ricardoErr "gitlab.com/ricardo-public/errors/pkg/errors"
	tokens "gitlab.com/ricardo-public/jwt-tools/v2/pkg/token"
	"gitlab.com/ricardo-public/tracing/pkg/tracing"
	authPort "gitlab.com/ricardo134/auth-service/internal/core/ports/auth"
	"gitlab.com/ricardo134/auth-service/internal/core/ports/user"
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
	notifier      user.EventsNotifier
}

func NewExternalTokenService(
	tRepo authPort.TokenRepository,
	aRepo authPort.AuthenticationRepository,
	notifier user.EventsNotifier,
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

func (e externalTokenService) Verify(ctx context.Context, token string) (*tokens.SignedTokens, error) {
	nctx, span := tracing.Tracer.Start(ctx, "auth.externalTokenService.Verify")
	defer span.End()

	user, err := e.tokenRepo.Verify(nctx, token)
	if err != nil {
		return nil, ricardoErr.New(ricardoErr.ErrUnauthorized, errors2.Wrap(err, "can't find user").Error())
	}

	existingUser, err := e.authRepo.EmailExists(nctx, user.Email)
	if existingUser == nil {
		existingUser, err = e.authRepo.Save(nctx, *user)
		if err != nil {
			return nil, err
		}

		_ = e.notifier.Created(*existingUser)
	}

	return generate(strconv.Itoa(int(existingUser.ID)), tokens.UserRole, e.accessSecret, e.refreshSecret), nil
}
