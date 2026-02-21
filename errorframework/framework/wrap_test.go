package framework

import (
	"errors"
	"testing"
)

func TestWrap(t *testing.T) {

	err := errors.New("original")

	appErr := Wrap(err, "wrapped")

	if appErr.Message != "wrapped" {
		t.Fatal("Wrap failed")
	}

	if appErr.Err == nil {
		t.Fatal("Original error not preserved")
	}
}

func TestWrapSafe(t *testing.T) {

	err := errors.New("original")

	appErr := WrapSafe(err, "wrapped")

	if appErr.IsSensitive {
		t.Fatal("WrapSafe should not be sensitive")
	}
}
