package boot

import (
	"gitlab.com/ricardo-public/tracing/pkg/tracing"
)

const (
	service = "auth-service"
)

func LoadTracing() {
	tracing.InitTracer(service, tracingEndpoint)
}
