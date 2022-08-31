package auth

import (
	"context"
	"gitlab.com/ricardo134/auth-service/internal/core/entities"
)

type TokenAuthenticate interface {
	Verify(ctx context.Context, token string) (*entities.SignedTokenPair, error)
}
