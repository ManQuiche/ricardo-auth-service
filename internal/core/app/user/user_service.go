package user

import (
	"github.com/pkg/errors"
	ricardoErr "gitlab.com/ricardo-public/errors/pkg/errors"
	"gitlab.com/ricardo134/auth-service/internal/core/entities"
	"gitlab.com/ricardo134/auth-service/internal/core/ports/user"
)

type Service interface {
	user.Service
}

type service struct {
	repo user.Repository
}

func NewUserService(r user.Repository) Service {
	return service{repo: r}
}

func (s service) Get(userID uint) (*entities.User, error) {
	user, err := s.repo.Get(userID)
	if err != nil {
		return nil, ricardoErr.New(ricardoErr.ErrNotFound, errors.Wrap(err, "cannot retrieve user").Error())
	}

	return user, err
}

func (s service) Update(user entities.User) (*entities.User, error) {
	updUser, err := s.repo.Update(user)
	if err != nil {
		return nil, ricardoErr.New(ricardoErr.ErrNotFound, errors.Wrap(err, "cannot update user").Error())
	}

	return updUser, err
}

func (s service) Delete(userID uint) (*entities.User, error) {
	delUser, err := s.repo.Delete(userID)
	if err != nil {
		return nil, ricardoErr.New(ricardoErr.ErrNotFound, errors.Wrap(err, "cannot delete user").Error())
	}

	return delUser, err
}
