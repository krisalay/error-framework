package framework

import (
	"net/http"

	"github.com/krisalay/error-framework/core"
)

// Wrap adds context to an existing error
func Wrap(err error, message string) *core.AppError {

	if err == nil {
		return nil
	}

	// If already AppError, wrap its internal error
	if appErr, ok := err.(*core.AppError); ok {

		return core.New().
			WithMessage(message).
			WithCode(appErr.Code).
			WithStatus(appErr.Status).
			WithDetails(appErr.Details).
			WithLevel(appErr.Level).
			WithSensitive(appErr.IsSensitive).
			WithInternal(appErr).
			Build()
	}

	// Wrap generic error
	return core.New().
		WithMessage(message).
		WithCode(core.CodeInternalError).
		WithStatus(http.StatusInternalServerError).
		WithLevel(core.LevelError).
		WithSensitive(true).
		WithInternal(err).
		Build()
}

func WrapWithCode(err error, code string, message string) *core.AppError {

	return core.New().
		WithMessage(message).
		WithCode(code).
		WithInternal(err).
		WithSensitive(true).
		Build()
}
