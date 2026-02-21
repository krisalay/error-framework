package framework

import (
	"context"

	"github.com/krisalay/error-framework/core"
)

// Recover converts panic into AppError and logs it
func Recover(ctx context.Context) *core.AppError {

	recovered := recover()

	if recovered == nil {
		return nil
	}

	f := get()

	appErr := core.New().
		WithMessage("Internal server error").
		WithCode(core.CodeInternalError).
		WithLevel(core.LevelFatal).
		WithSensitive(true).
		WithDetail("panic", recovered).
		Build()

	f.manager.Handle(ctx, appErr)

	return appErr
}

func RecoverAndWrap(ctx context.Context, message string) *core.AppError {

	recovered := recover()

	if recovered == nil {
		return nil
	}

	return Wrap(
		core.New().
			WithDetail("panic", recovered).
			Build(),
		message,
	)
}
