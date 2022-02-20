package auth

import (
	"context"
	"ricardo/auth-service/internal/core/entities"
)

type AuthenticationRepository interface {
	Exists(ctx context.Context, username, password string) (*entities.User, error)
	Save(ctx context.Context, user entities.User) error
}
