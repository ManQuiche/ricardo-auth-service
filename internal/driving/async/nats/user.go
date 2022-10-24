package nats

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"gitlab.com/ricardo-public/tracing/pkg/tracing"
	"gitlab.com/ricardo134/auth-service/internal/core/ports/user"
	"go.opentelemetry.io/otel/trace"
	"log"
)

type UserHandler interface {
	Requested(awt tracing.AnyWithTrace)
}

type userHandler struct {
	userService user.Service
}

func NewNatsUserHandler(userSvc user.Service) UserHandler {
	return userHandler{userSvc}
}

func (nh userHandler) Requested(awt tracing.AnyWithTrace) {
	traceID, err := trace.TraceIDFromHex(awt.TraceID)
	if err != nil {
		log.Println(errors.Wrap(err, fmt.Sprintf("cannot parse traceID %s", awt.TraceID)).Error())
	}

	ctx := trace.ContextWithRemoteSpanContext(context.Background(), trace.NewSpanContext(
		trace.SpanContextConfig{
			TraceID: traceID,
		},
	))
	nctx, span := tracing.Tracer.Start(ctx, "nats.UserHandler.Requested")
	defer span.End()

	userID, ok := awt.Any.(uint)
	if ok == false {

	}

	_, _ = nh.userService.Get(nctx, userID)
}
