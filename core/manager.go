package core

import (
	"context"
	"errors"
	"time"
)

type Manager struct {
	logger             Logger
	traceProvider      TraceProvider
	stackTraceProvider StackTraceProvider
}

// ManagerConfig allows flexible initialization
type ManagerConfig struct {
	Logger             Logger
	TraceProvider      TraceProvider
	StackTraceProvider StackTraceProvider
}

// NewManager creates a new error manager
func NewManager(config ManagerConfig) *Manager {

	if config.Logger == nil {
		panic("Logger is required for error manager")
	}

	return &Manager{
		logger:             config.Logger,
		traceProvider:      config.TraceProvider,
		stackTraceProvider: config.StackTraceProvider,
	}
}

func (m *Manager) Handle(ctx context.Context, err error) *AppError {

	if err == nil {
		return nil
	}

	// If already AppError, enrich and return
	var appErr *AppError
	if errors.As(err, &appErr) {

		m.enrich(ctx, appErr)
		m.logger.Log(appErr)

		return appErr
	}

	// Unknown error → convert to internal error
	appErr = New().
		WithMessage("Internal server error").
		WithCode(CodeInternalError).
		WithStatus(500).
		WithLevel(LevelError).
		WithSensitive(true).
		WithInternal(err).
		Build()

	m.enrich(ctx, appErr)
	m.logger.Log(appErr)

	return appErr
}

func (m *Manager) enrich(ctx context.Context, err *AppError) {

	if err.Timestamp.IsZero() {
		err.Timestamp = time.Now()
	}

	if m.traceProvider != nil && err.TraceID == "" {
		err.TraceID = m.traceProvider.GetTraceID(ctx)
	}

	if m.stackTraceProvider != nil && err.StackTrace == "" {
		err.StackTrace = m.stackTraceProvider.Capture()
	}
}

func (m *Manager) ToResponse(ctx context.Context, err error) *AppError {

	appErr := m.Handle(ctx, err)

	// Return sanitized copy
	return &AppError{
		Message:   appErr.SafeMessage(),
		Code:      appErr.SafeCode(),
		Status:    appErr.Status,
		TraceID:   appErr.TraceID,
		Timestamp: appErr.Timestamp,
	}
}

func (m *Manager) HandlePanic(ctx context.Context, recovered any) *AppError {

	appErr := New().
		WithMessage("Internal server error").
		WithCode(CodeInternalError).
		WithStatus(500).
		WithLevel(LevelFatal).
		WithSensitive(true).
		WithDetail("panic", recovered).
		Build()

	m.enrich(ctx, appErr)
	m.logger.Log(appErr)

	return appErr
}

func (m *Manager) Wrap(ctx context.Context, err error, message string) *AppError {

	if err == nil {
		return nil
	}

	appErr := New().
		WithMessage(message).
		WithCode(CodeInternalError).
		WithStatus(500).
		WithSensitive(true).
		WithInternal(err).
		Build()

	m.enrich(ctx, appErr)
	m.logger.Log(appErr)

	return appErr
}
