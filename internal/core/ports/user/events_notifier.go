package user

import (
	"context"
	"gitlab.com/ricardo134/auth-service/internal/core/entities"
)

// EventsNotifier Interface designed as a template for event
// publishing function on user registration
type EventsNotifier interface {
	Created(ctx context.Context, user entities.User) error
	Updated(ctx context.Context, user entities.User) error
	Deleted(ctx context.Context, userID uint) error
}
