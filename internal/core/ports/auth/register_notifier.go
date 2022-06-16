package auth

import "ricardo/auth-service/internal/core/entities"

type RegisterNotifier interface {
	Notify(user entities.User) error
}
