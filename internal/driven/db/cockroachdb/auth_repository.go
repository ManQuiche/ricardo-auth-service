package cockroachdb

import (
	"auth-service/internal/core/entities"
	"auth-service/internal/core/ports/auth"
	"context"
	"errors"
	"gorm.io/gorm"
)

type authenticationRepository struct {
	client *gorm.DB
}

func NewAuthenticationRepository(client *gorm.DB) auth.AuthenticationRepository {
	return authenticationRepository{
		client: client,
	}
}

func (r authenticationRepository) Exists(ctx context.Context, username, password string) (*entities.User, error) {
	var user *entities.User
	r.client.Where("username = ? and password = ?", username, password).First(user)

	if user == nil {
		return nil, errors.New("cannot find user")
	}

	return user, nil
}

func (r authenticationRepository) Save(ctx context.Context, user entities.User) error {
	r.client.Save(user)

	return nil
}
