package user

import "gitlab.com/ricardo134/auth-service/internal/core/entities"

// EventsNotifier Interface designed as a template for event
// publishing function on user registration
type EventsNotifier interface {
	Created(user entities.User) error
	Updated(user entities.User) error
	Deleted(userID uint) error
}
