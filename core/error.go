package core

import (
	"time"
)

type AppError struct {
	Message string
	Code    string
	Status  int
	Details map[string]any

	Level       ErrorLevel
	Err         error
	IsSensitive bool

	Timestamp  time.Time
	StackTrace string
	TraceID    string
}

// Implements Go error interface
func (e *AppError) Error() string {
	return e.Message
}

// Unwrap allows errors.Unwrap support
func (e *AppError) Unwrap() error {
	return e.Err
}

// SafeMessage returns client-safe message
func (e *AppError) SafeMessage() string {
	if e.IsSensitive {
		return "Internal server error"
	}
	return e.Message
}

// SafeCode returns client-safe code
func (e *AppError) SafeCode() string {
	if e.IsSensitive {
		return CodeInternalError
	}
	return e.Code
}
