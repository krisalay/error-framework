package core

import (
	"errors"
	"testing"
)

func TestAppError_Error(t *testing.T) {

	err := &AppError{
		Message: "test error",
	}

	if err.Error() != "test error" {
		t.Fatal("Error() did not return correct message")
	}
}

func TestAppError_Unwrap(t *testing.T) {

	original := errors.New("original")

	appErr := &AppError{
		Err: original,
	}

	if appErr.Unwrap() != original {
		t.Fatal("Unwrap failed")
	}
}

func TestSafeMessage_Sensitive(t *testing.T) {

	appErr := &AppError{
		Message:     "secret",
		IsSensitive: true,
	}

	if appErr.SafeMessage() != "Internal server error" {
		t.Fatal("Sensitive message leaked")
	}
}

func TestSafeMessage_NotSensitive(t *testing.T) {

	appErr := &AppError{
		Message:     "visible",
		IsSensitive: false,
	}

	if appErr.SafeMessage() != "visible" {
		t.Fatal("Safe message incorrect")
	}
}
