package utils

import (
	"context"
	"testing"
)

func TestTraceProvider(t *testing.T) {

	provider := NewTraceProvider()

	ctx := context.Background()

	trace := provider.GetTraceID(ctx)

	if trace == "" {
		t.Fatal("Trace ID empty")
	}
}
