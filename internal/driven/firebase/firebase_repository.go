package firebase

import (
	"context"
	"gitlab.com/ricardo134/auth-service/internal/core/entities"
	"log"

	"gitlab.com/ricardo134/auth-service/internal/core/ports/auth"

	fireAuth "firebase.google.com/go/auth"
)

type TokenRepository interface {
	auth.TokenRepository
}

const userSource = "firebase"

type tokenRepository struct {
	firebaseClient *fireAuth.Client
}

func NewTokenRepository(firebaseClient *fireAuth.Client) TokenRepository {
	return tokenRepository{firebaseClient: firebaseClient}
}

func (t tokenRepository) Verify(ctx context.Context, token string) (*entities.User, error) {
	fToken, err := t.firebaseClient.VerifyIDTokenAndCheckRevoked(ctx, token)

	if err != nil {
		return nil, err
	}

	// FIXME: remove this when done testing
	log.Println("retrieved token from Firebase :")
	log.Println(fToken)

	return &entities.User{
		// TODO: check if Subject is really an email
		Email:          fToken.Subject,
		ExternalSource: userSource,
	}, nil
}
