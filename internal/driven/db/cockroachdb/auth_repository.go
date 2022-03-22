package cockroachdb

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"ricardo/auth-service/internal/core/entities"
	"ricardo/auth-service/internal/core/ports/auth"
)

type authenticationRepository struct {
	client *gorm.DB
}

func NewAuthenticationRepository(client *gorm.DB) auth.AuthenticationRepository {
	return authenticationRepository{
		client: client,
	}
}

func (r authenticationRepository) Exists(ctx context.Context, email, password string) (*entities.User, error) {
	var user *entities.User
	r.client.Where("email = ? and password = ?", email, password).First(&user)

	if user == nil {
		return nil, errors.New("cannot find user")
	}

	return user, nil
}

func (r authenticationRepository) Save(ctx context.Context, user entities.User) error {
	r.client.Save(&user)

	return nil
}
