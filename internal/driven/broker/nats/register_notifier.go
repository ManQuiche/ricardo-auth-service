package nats

import (
	"github.com/nats-io/nats.go"
	"ricardo/auth-service/internal/core/entities"
	"ricardo/auth-service/internal/core/ports/auth"
)

type registerNotifier struct {
	conn  *nats.EncodedConn
	topic string
}

func NewRegisterNotifier(conn *nats.EncodedConn, topic string) auth.RegisterNotifier {
	return registerNotifier{
		conn:  conn,
		topic: topic,
	}
}

func (r registerNotifier) Notify(user entities.User) error {
	return r.conn.Publish(r.topic, &user)
}
