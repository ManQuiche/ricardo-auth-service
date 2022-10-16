package user

import (
	"context"
	"gitlab.com/ricardo134/auth-service/internal/core/entities"
)

type Service interface {
	Get(ctx context.Context, userID uint) (*entities.User, error)
	Update(ctx context.Context, user entities.User) (*entities.User, error)
	Delete(ctx context.Context, userID uint) (*entities.User, error)
}

type Repository interface {
	Service
}
