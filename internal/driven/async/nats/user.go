package nats

import (
	"context"
	"github.com/nats-io/nats.go"
	"gitlab.com/ricardo-public/tracing/pkg/tracing"
	"gitlab.com/ricardo134/auth-service/internal/core/entities"
	userports "gitlab.com/ricardo134/auth-service/internal/core/ports/user"
)

type userEventsNotifier struct {
	conn         *nats.EncodedConn
	createdTopic string
	updatedTopic string
	deletedTopic string
}

func NewUserEventsNotifier(conn *nats.EncodedConn, createdTopic, updatedTopic, deletedTopic string) userports.EventsNotifier {
	return userEventsNotifier{conn, createdTopic, updatedTopic, deletedTopic}
}

func (r userEventsNotifier) Created(ctx context.Context, user entities.User) error {
	_, span := tracing.Tracer.Start(ctx, "nats.userEventsNotifier.Created")
	defer span.End()

	withTrace := tracing.AnyWithTrace[entities.ShortUser]{
		Any: entities.ShortUser{
			ID:       user.ID,
			Username: user.Username,
		},
		TraceID: span.SpanContext().TraceID().String(),
	}

	return r.conn.Publish(r.createdTopic, &withTrace)
}

func (r userEventsNotifier) Updated(ctx context.Context, user entities.User) error {
	_, span := tracing.Tracer.Start(ctx, "nats.userEventsNotifier.Updated")
	defer span.End()

	withTrace := tracing.AnyWithTrace[entities.ShortUser]{
		Any: entities.ShortUser{
			ID:       user.ID,
			Username: user.Username,
		},
		TraceID: span.SpanContext().TraceID().String(),
	}
	return r.conn.Publish(r.updatedTopic, &withTrace)
}

func (r userEventsNotifier) Deleted(ctx context.Context, userID uint) error {
	_, span := tracing.Tracer.Start(ctx, "nats.userEventsNotifier.Deleted")
	defer span.End()

	withTrace := tracing.AnyWithTrace[uint]{
		Any:     userID,
		TraceID: span.SpanContext().TraceID().String(),
	}
	return r.conn.Publish(r.deletedTopic, &withTrace)
}
