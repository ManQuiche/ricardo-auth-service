package auth

import (
	"context"
	"ricardo/auth-service/internal/core/entities"
)

type TokenRepository interface {
	// Verify Please return a User whom field ExternalSource is defined by your source name (ex: firebase)
	Verify(ctx context.Context, token string) (*entities.User, error)
}
