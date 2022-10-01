package user

import (
	"gitlab.com/ricardo134/auth-service/internal/core/entities"
)

type Service interface {
	Get(userID uint) (*entities.User, error)
	Update(user entities.User) (*entities.User, error)
	Delete(userID uint) (*entities.User, error)
}

type Repository interface {
	Service
}
