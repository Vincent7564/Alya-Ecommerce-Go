package tracing

import (
	"context"

	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

// InitializeTracerProvider sets up the OpenTelemetry tracer provider
func InitializeTracerProvider(serviceName string) (*sdktrace.TracerProvider, error) {
	ctx := context.Background()

	// Configure OTLP exporter
	client := otlptracegrpc.NewClient(
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint("localhost:4317"), // Or your Jaeger endpoint
	)
	exporter, err := otlptrace.New(ctx, client)
	if err != nil {
		log.Error().Err(err).Msg("Cannot Connect to Client")
		return nil, err
	}

	res, err := resource.New(ctx,
		resource.WithSchemaURL(semconv.SchemaURL), // Use the desired version
		resource.WithAttributes(
			semconv.ServiceName(serviceName),
			attribute.String("environment", "development"),
		),
	)
	if err != nil {
		log.Error().Err(err).Msg("Cannot Connect to Client")
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	return tp, nil
}
