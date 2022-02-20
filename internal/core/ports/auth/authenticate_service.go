package auth

import (
	"context"
	"ricardo/auth-service/internal/core/entities"
	"ricardo/auth-service/internal/core/entities/auth"
)

type Authenticate interface {
	Login(ctx context.Context, loginRequest entities.LoginRequest) (*auth.SignedTokenPair, error)
	Save(ctx context.Context, user entities.User) error
}
