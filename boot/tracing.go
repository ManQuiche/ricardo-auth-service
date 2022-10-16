package boot

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"go.opentelemetry.io/otel/trace"
	"log"
)

const (
	service = "auth-service"
)

var (
	Tracer trace.Tracer
)

func LoadTracing() {
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(tracingEndpoint)))
	if err != nil {
		log.Fatalf("cannot instanciate jaeger endpoint %s : %s", tracingEndpoint, err)
	}

	tp := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exporter),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(service),
			semconv.TelemetrySDKLanguageGo,
		)),
	)

	otel.SetTracerProvider(tp)
	Tracer = otel.Tracer(service)
}
