package auth

import (
	"auth-service/internal/core/entities"
	"auth-service/internal/core/entities/auth"
	"context"
)

type Authenticate interface {
	Login(ctx context.Context, loginRequest entities.LoginRequest) (*auth.SignedTokenPair, error)
	Save(ctx context.Context, user entities.User) error
}
