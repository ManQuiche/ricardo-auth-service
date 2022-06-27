package auth

import "ricardo/auth-service/internal/core/entities"

// RegisterNotifier Interface designed as a template for event
// publishing function on user registration
type RegisterNotifier interface {
	Notify(user entities.User) error
}
