package utils

import "testing"

func TestStackTrace(t *testing.T) {

	provider := NewStackTraceProvider()

	trace := provider.Capture()

	if trace == "" {
		t.Fatal("Stacktrace empty")
	}
}
