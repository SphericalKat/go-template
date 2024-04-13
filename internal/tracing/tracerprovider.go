package tracing

import (
	"context"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

var tp *sdktrace.TracerProvider

// ProvideTracerProvider provides a tracer provider
func ProvideTracerProvider() (*sdktrace.TracerProvider, error) {
	if tp == nil {
		ctx := context.Background()
		// ─── Configure Tracer Provider ────────────────────────────────────────
		exporter, err := otlptracehttp.New(ctx, otlptracehttp.WithEndpoint("localhost:4318"), otlptracehttp.WithInsecure())
		if err != nil {
			return nil, err
		}

		resource, err := resource.New(ctx,
			resource.WithAttributes(
				semconv.ServiceNameKey.String("setsu"),
				semconv.ServiceVersionKey.String("0.1.0"),
				attribute.String("environment", "local"),
			),
		)
		if err != nil {
			return nil, err
		}

		tp = sdktrace.NewTracerProvider(
			sdktrace.WithBatcher(exporter),
			sdktrace.WithResource(resource),
		)
	}

	return tp, nil
}
