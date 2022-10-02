package nats

import (
	"github.com/nats-io/nats.go"
	"gitlab.com/ricardo134/auth-service/internal/core/entities"
	userports "gitlab.com/ricardo134/auth-service/internal/core/ports/user"
)

type eventsNotifier struct {
	conn         *nats.EncodedConn
	createdTopic string
	updatedTopic string
	deletedTopic string
}

func NewUserEventsNotifier(conn *nats.EncodedConn, createdTopic, updatedTopic, deletedTopic string) userports.EventsNotifier {
	return eventsNotifier{conn, createdTopic, updatedTopic, deletedTopic}
}

func (r eventsNotifier) Created(user entities.User) error {
	shortUser := entities.ShortUser{
		ID:       user.ID,
		Username: user.Username,
	}
	return r.conn.Publish(r.createdTopic, &shortUser)
}

func (r eventsNotifier) Updated(user entities.User) error {
	shortUser := entities.ShortUser{
		ID:       user.ID,
		Username: user.Username,
	}
	return r.conn.Publish(r.updatedTopic, &shortUser)
}

func (r eventsNotifier) Deleted(userID uint) error {
	return r.conn.Publish(r.deletedTopic, userID)
}
