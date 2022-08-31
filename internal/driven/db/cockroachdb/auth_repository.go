package cockroachdb

import (
	"context"
	"errors"
	ricardoerr "gitlab.com/ricardo-public/errors/pkg/errors"
	"gitlab.com/ricardo134/auth-service/internal/core/entities"
	"gitlab.com/ricardo134/auth-service/internal/core/ports/auth"
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

func notFoundOrElseError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ricardoerr.New(ricardoerr.ErrNotFound, "record not found")
	}

	return ricardoerr.New(ricardoerr.ErrDatabaseError, err.Error())
}

func (r authenticationRepository) Exists(ctx context.Context, email, password string) (*entities.User, error) {
	var user *entities.User
	err := r.client.Where("email = ? and password = ?", email, password).First(&user).Error
	if err != nil {
		return nil, notFoundOrElseError(err)
	}

	return user, nil
}

func (r authenticationRepository) Save(ctx context.Context, user entities.User) (*entities.User, error) {
	err := r.client.Save(&user).Error
	if err != nil {
		return nil, notFoundOrElseError(err)
	}

	return &user, nil
}

func (r authenticationRepository) EmailExists(ctx context.Context, email string) (*entities.User, error) {
	var user *entities.User
	err := r.client.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, notFoundOrElseError(err)
	}

	return user, nil
}
