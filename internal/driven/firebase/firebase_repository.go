package firebase

import (
	"context"
	"gitlab.com/ricardo134/auth-service/internal/core/entities"
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
	fToken, err := t.firebaseClient.VerifyIDToken(ctx, token)

	if err != nil {
		return nil, err
	}

	return &entities.User{
		Email:          fToken.Claims["email"].(string),
		Username:       fToken.Claims["name"].(string),
		ExternalSource: userSource,
	}, nil
}
