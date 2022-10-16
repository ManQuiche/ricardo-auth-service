package user

import (
	"context"
	"github.com/pkg/errors"
	errorsext "gitlab.com/ricardo-public/errors/pkg/errors"
	"gitlab.com/ricardo-public/tracing/pkg/tracing"
	"gitlab.com/ricardo134/auth-service/internal/core/entities"
	userports "gitlab.com/ricardo134/auth-service/internal/core/ports/user"
)

type Service interface {
	userports.Service
}

type service struct {
	repo     userports.Repository
	notifier userports.EventsNotifier
}

func NewService(r userports.Repository, notifier userports.EventsNotifier) Service {
	return service{r, notifier}
}

func (s service) Get(ctx context.Context, userID uint) (*entities.User, error) {
	nctx, span := tracing.Tracer.Start(ctx, "user.service.Get")
	var err error
	defer span.End()

	user, err := s.repo.Get(nctx, userID)
	if err != nil {
		return nil, errorsext.New(errorsext.ErrNotFound, errors.Wrap(err, "cannot retrieve user").Error())
	}

	return user, err
}

func (s service) Update(ctx context.Context, user entities.User) (*entities.User, error) {
	nctx, span := tracing.Tracer.Start(ctx, "user.service.Update")
	var err error
	defer span.End()

	updUser, err := s.repo.Update(nctx, user)
	if err != nil {
		return nil, errorsext.New(errorsext.ErrNotFound, errors.Wrap(err, "cannot update user").Error())
	}

	_ = s.notifier.Updated(user)

	return updUser, err
}

func (s service) Delete(ctx context.Context, userID uint) (*entities.User, error) {
	nctx, span := tracing.Tracer.Start(ctx, "user.service.Delete")
	var err error
	defer span.End()

	delUser, err := s.repo.Delete(nctx, userID)
	if err != nil {
		return nil, errorsext.New(errorsext.ErrNotFound, errors.Wrap(err, "cannot delete user").Error())
	}

	_ = s.notifier.Deleted(userID)

	return delUser, err
}
