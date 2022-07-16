package auth

import (
	"context"
	"ricardo/auth-service/internal/core/entities"
)

type TokenAuthenticate interface {
	Verify(ctx context.Context, token string) (*entities.SignedTokenPair, error)
}
