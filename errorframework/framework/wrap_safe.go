package framework

import (
	"net/http"

	"github.com/krisalay/error-framework/core"
)

// WrapSafe wraps an error with a safe message that can be exposed to clients.
func WrapSafe(err error, message string) *core.AppError {

	if err == nil {
		return nil
	}

	// If already AppError, preserve properties but mark safe
	if appErr, ok := err.(*core.AppError); ok {

		return core.New().
			WithMessage(message).
			WithCode(appErr.Code).
			WithStatus(appErr.Status).
			WithDetails(appErr.Details).
			WithLevel(appErr.Level).
			WithSensitive(false).
			WithInternal(appErr).
			Build()
	}

	// Generic error wrapping
	return core.New().
		WithMessage(message).
		WithCode(core.CodeInternalError).
		WithStatus(http.StatusInternalServerError).
		WithLevel(core.LevelError).
		WithSensitive(false).
		WithInternal(err).
		Build()
}
