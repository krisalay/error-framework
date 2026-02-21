package utils

import (
	"fmt"
	"runtime"
	"strings"
)

type StackTraceProvider struct {
	skipFrames int
}

func NewStackTraceProvider() *StackTraceProvider {
	return &StackTraceProvider{
		skipFrames: 3, // skip runtime + this file + manager
	}
}

// Capture returns formatted stack trace string
func (s *StackTraceProvider) Capture() string {

	const maxDepth = 32

	pcs := make([]uintptr, maxDepth)

	n := runtime.Callers(s.skipFrames, pcs)

	frames := runtime.CallersFrames(pcs[:n])

	var builder strings.Builder

	for {
		frame, more := frames.Next()

		// Skip runtime frames
		if !isFrameworkFrame(frame.Function) {

			builder.WriteString(fmt.Sprintf(
				"%s\n\t%s:%d\n",
				frame.Function,
				frame.File,
				frame.Line,
			))
		}

		if !more {
			break
		}
	}

	return builder.String()
}

func isFrameworkFrame(function string) bool {

	if strings.Contains(function, "runtime.") {
		return true
	}

	if strings.Contains(function, "errorframework") {
		return true
	}

	return false
}
