package framework

// import (
// 	"context"

// 	"github.com/krisalay/error-framework/core"
// )

// func Recover(ctx context.Context) *core.AppError {

// 	recovered := recover()

// 	if recovered == nil {
// 		return nil
// 	}

// 	appErr := core.New().
// 		WithMessage("Internal server error").
// 		WithCode(core.CodeInternalError).
// 		WithLevel(core.LevelFatal).
// 		WithSensitive(true).
// 		WithDetail("panic", recovered).
// 		Build()

// 	// SAFE access to instance (DO NOT call get())
// 	if instance != nil && instance.manager != nil {
// 		instance.manager.Handle(ctx, appErr)
// 	}

// 	return appErr
// }

// func RecoverAndWrap(ctx context.Context, message string) *core.AppError {

// 	recovered := recover()

// 	if recovered == nil {
// 		return nil
// 	}

// 	return Wrap(
// 		core.New().
// 			WithDetail("panic", recovered).
// 			Build(),
// 		message,
// 	)
// }
