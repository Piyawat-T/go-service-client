package bootstrap

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

func NewTraceProvider() *sdktrace.TracerProvider {
	tp := sdktrace.NewTracerProvider()
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.TraceContext{})
	return tp
}

type ctxKey struct{}

func NewContext(parent context.Context, t trace.Tracer) context.Context {
	return context.WithValue(parent, ctxKey{}, t)
}
