package utils

import (
	"context"

	"github.com/google/uuid"
)

type contextKey string

const TraceIDKey contextKey = "trace_id"

type TraceProvider struct{}

func NewTraceProvider() *TraceProvider {
	return &TraceProvider{}
}

// GetTraceID returns trace id from context or generates new one
func (t *TraceProvider) GetTraceID(ctx context.Context) string {

	if ctx == nil {
		return t.generate()
	}

	traceID, ok := ctx.Value(TraceIDKey).(string)

	if ok && traceID != "" {
		return traceID
	}

	return t.generate()
}

func (t *TraceProvider) generate() string {
	return uuid.New().String()
}

// Inject adds trace id to context
func (t *TraceProvider) Inject(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, TraceIDKey, traceID)
}
