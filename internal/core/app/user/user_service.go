package user

import (
	"github.com/pkg/errors"
	errorsext "gitlab.com/ricardo-public/errors/pkg/errors"
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

func (s service) Get(userID uint) (*entities.User, error) {
	user, err := s.repo.Get(userID)
	if err != nil {
		return nil, errorsext.New(errorsext.ErrNotFound, errors.Wrap(err, "cannot retrieve user").Error())
	}

	return user, err
}

func (s service) Update(user entities.User) (*entities.User, error) {
	updUser, err := s.repo.Update(user)
	if err != nil {
		return nil, errorsext.New(errorsext.ErrNotFound, errors.Wrap(err, "cannot update user").Error())
	}

	return updUser, err
}

func (s service) Delete(userID uint) (*entities.User, error) {
	delUser, err := s.repo.Delete(userID)
	if err != nil {
		return nil, errorsext.New(errorsext.ErrNotFound, errors.Wrap(err, "cannot delete user").Error())
	}

	_ = s.notifier.Deleted(userID)

	return delUser, err
}
