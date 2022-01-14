package auth

import (
	"auth-service/internal/core/entities"
	"auth-service/internal/core/entities/auth"
	"context"
)

type Authenticate interface {
	Login(ctx context.Context, username, password string) (*auth.TokenPair, error)
	Save(ctx context.Context, user entities.User) error
}
