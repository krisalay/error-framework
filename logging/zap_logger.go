package logging

import (
	"os"

	"github.com/krisalay/error-framework/core"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	logger *zap.Logger
}

func NewZapLogger(config Config) (*ZapLogger, error) {

	level := parseLevel(config.Level)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	var encoder zapcore.Encoder

	if config.Encoding == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	var cores []zapcore.Core

	// Console logging
	if config.ConsoleEnabled {
		consoleCore := zapcore.NewCore(
			encoder,
			zapcore.AddSync(os.Stdout),
			level,
		)
		cores = append(cores, consoleCore)
	}

	// File logging
	if config.FileEnabled {

		file, err := os.OpenFile(
			config.FilePath,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0644,
		)

		if err != nil {
			return nil, err
		}

		fileCore := zapcore.NewCore(
			encoder,
			zapcore.AddSync(file),
			level,
		)

		cores = append(cores, fileCore)
	}

	coreCombined := zapcore.NewTee(cores...)

	logger := zap.New(
		coreCombined,
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)

	return &ZapLogger{
		logger: logger,
	}, nil
}

func (z *ZapLogger) Log(err *core.AppError) {

	fields := []zap.Field{
		zap.String("code", err.Code),
		zap.Int("status", err.Status),
		zap.String("level", err.Level.String()),
		zap.String("trace_id", err.TraceID),
		zap.Time("timestamp", err.Timestamp),
	}

	if err.StackTrace != "" {
		fields = append(fields,
			zap.String("stacktrace", err.StackTrace))
	}

	if err.Err != nil {
		fields = append(fields,
			zap.Error(err.Err))
	}

	if err.Details != nil {
		fields = append(fields,
			zap.Any("details", err.Details))
	}

	switch err.Level {

	case core.LevelDebug:
		z.logger.Debug(err.Message, fields...)

	case core.LevelInfo:
		z.logger.Info(err.Message, fields...)

	case core.LevelWarn:
		z.logger.Warn(err.Message, fields...)

	case core.LevelError:
		z.logger.Error(err.Message, fields...)

	case core.LevelFatal:
		z.logger.Fatal(err.Message, fields...)

	default:
		z.logger.Error(err.Message, fields...)
	}
}
