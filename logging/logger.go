package logging

import "go.uber.org/zap/zapcore"

type Config struct {
	ConsoleEnabled bool
	FileEnabled    bool
	FilePath       string
	Level          string // debug, info, warn, error, fatal
	Encoding       string // json or console
}

func parseLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.ErrorLevel
	}
}
