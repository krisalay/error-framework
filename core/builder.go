package core

import (
	"net/http"
	"time"
)

type Builder struct {
	err *AppError
}

func New() *Builder {
	return &Builder{
		err: &AppError{
			Message:     "Internal server error",
			Code:        CodeInternalError,
			Status:      http.StatusInternalServerError,
			Level:       LevelError,
			IsSensitive: true,
			Timestamp:   time.Now(),
			Details:     make(map[string]any),
		},
	}
}

func (b *Builder) WithMessage(message string) *Builder {
	b.err.Message = message
	return b
}

func (b *Builder) WithCode(code string) *Builder {
	b.err.Code = code
	return b
}

func (b *Builder) WithStatus(status int) *Builder {
	b.err.Status = status
	return b
}

func (b *Builder) WithDetails(details map[string]any) *Builder {
	b.err.Details = details
	return b
}

func (b *Builder) WithDetail(key string, value any) *Builder {
	b.err.Details[key] = value
	return b
}

func (b *Builder) WithLevel(level ErrorLevel) *Builder {
	b.err.Level = level
	return b
}

func (b *Builder) WithInternal(err error) *Builder {
	b.err.Err = err
	return b
}

func (b *Builder) WithSensitive(sensitive bool) *Builder {
	b.err.IsSensitive = sensitive
	return b
}

func (b *Builder) WithTraceID(traceID string) *Builder {
	b.err.TraceID = traceID
	return b
}

func (b *Builder) WithStackTrace(stack string) *Builder {
	b.err.StackTrace = stack
	return b
}

func (b *Builder) Build() *AppError {
	return b.err
}
