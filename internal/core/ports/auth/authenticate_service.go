package auth

import (
	"context"
	"ricardo/auth-service/internal/core/entities"
)

type Authenticate interface {
	Login(ctx context.Context, loginRequest entities.LoginRequest) (*entities.SignedTokenPair, error)
	Save(ctx context.Context, user entities.User) error
	Refresh(ctx context.Context, token string) (*entities.SignedTokenPair, error)
}
