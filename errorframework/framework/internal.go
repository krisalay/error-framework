package framework

import (
	"net/http"

	"github.com/krisalay/error-framework/core"
)

func Internal(err error) *core.AppError {

	if err == nil {
		return nil
	}

	return core.New().
		WithMessage("Internal server error").
		WithCode(core.CodeInternalError).
		WithStatus(http.StatusInternalServerError).
		WithSensitive(true).
		WithInternal(err).
		WithLevel(core.LevelError).
		Build()
}

func NotFound(message string) *core.AppError {

	return core.New().
		WithMessage(message).
		WithCode(core.CodeNotFound).
		WithStatus(404).
		WithSensitive(false).
		WithLevel(core.LevelInfo).
		Build()
}

func AlreadyExists(message string) *core.AppError {

	return core.New().
		WithMessage(message).
		WithCode(core.CodeAlreadyExists).
		WithStatus(409).
		WithSensitive(false).
		WithLevel(core.LevelWarn).
		Build()
}
