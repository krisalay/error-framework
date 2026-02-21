package core

import "context"

type Logger interface {
	Log(err *AppError)
}

type TraceProvider interface {
	GetTraceID(ctx context.Context) string
}

type StackTraceProvider interface {
	Capture() string
}
