package auth

import "auth-service/internal/core/entities"

type AuthenticationRepository interface {
	Exists(username, password string) (*entities.User, error)
}
