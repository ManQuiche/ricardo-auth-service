package postgresql

import (
	"context"
	"gitlab.com/ricardo-public/tracing/pkg/tracing"
	"gitlab.com/ricardo134/auth-service/internal/core/entities"
	"gitlab.com/ricardo134/auth-service/internal/core/ports/auth"
	"go.opentelemetry.io/otel/codes"
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

func (r authenticationRepository) Exists(ctx context.Context, email, password string) (*entities.User, error) {
	_, span := tracing.Tracer.Start(ctx, "postgres.authenticationRepository.Exists")
	defer span.End()

	var user *entities.User
	err := r.client.Where("email = ? and password = ?", email, password).First(&user).Error
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, notFoundOrElseError(err)
	}

	return user, nil
}

func (r authenticationRepository) Save(ctx context.Context, user entities.User) (*entities.User, error) {
	_, span := tracing.Tracer.Start(ctx, "postgres.authenticationRepository.Save")
	defer span.End()

	err := r.client.Save(&user).Error
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, notFoundOrElseError(err)
	}

	return &user, nil
}

func (r authenticationRepository) EmailExists(ctx context.Context, email string) (*entities.User, error) {
	_, span := tracing.Tracer.Start(ctx, "postgres.authenticationRepository.EmailExists")
	defer span.End()

	var user *entities.User
	err := r.client.Where("email = ?", email).First(&user).Error
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, notFoundOrElseError(err)
	}

	return user, nil
}
