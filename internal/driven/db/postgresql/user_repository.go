package postgresql

import (
	"context"
	"gitlab.com/ricardo-public/tracing/pkg/tracing"
	"gitlab.com/ricardo134/auth-service/internal/core/entities"
	"gitlab.com/ricardo134/auth-service/internal/core/ports/user"
	"gorm.io/gorm"
)

type userRepository struct {
	client *gorm.DB
}

func NewUserRepository(client *gorm.DB) user.Repository {
	return userRepository{
		client: client,
	}
}

func (u userRepository) Get(ctx context.Context, userID uint) (*entities.User, error) {
	_, span := tracing.Tracer.Start(ctx, "postgresql.userRepository.Get")
	defer span.End()

	var gUser *entities.User
	err := u.client.First(&gUser, userID).Error
	if err != nil {
		return nil, notFoundOrElseError(err)
	}

	return gUser, nil
}

func (u userRepository) Update(ctx context.Context, user entities.User) (*entities.User, error) {
	_, span := tracing.Tracer.Start(ctx, "postgresql.userRepository.Update")
	defer span.End()

	err := u.client.Save(&user).Error
	if err != nil {
		return nil, notFoundOrElseError(err)
	}

	return &user, nil
}

func (u userRepository) Delete(ctx context.Context, userID uint) (*entities.User, error) {
	_, span := tracing.Tracer.Start(ctx, "postgresql.userRepository.Delete")
	defer span.End()

	var delUser *entities.User
	err := u.client.Delete(&delUser, userID).Error
	if err != nil {
		return nil, notFoundOrElseError(err)
	}

	return delUser, nil
}
