package framework

import (
	pgxadapter "github.com/krisalay/error-framework/adapters/pgx"
	validatoradapter "github.com/krisalay/error-framework/adapters/validator"
	"github.com/krisalay/error-framework/core"
	"github.com/krisalay/error-framework/errorframework/config"
	"github.com/krisalay/error-framework/logging"
	"github.com/krisalay/error-framework/utils"
)

func InitFromConfig(cfg config.Config) (*core.Manager, error) {

	// Logger
	logger, err := logging.NewZapLogger(logging.Config{
		ConsoleEnabled: cfg.Logger.ConsoleEnabled,
		FileEnabled:    cfg.Logger.FileEnabled,
		FilePath:       cfg.Logger.FilePath,
		Level:          cfg.Logger.Level,
		Encoding:       cfg.Logger.Encoding,
	})

	if err != nil {
		return nil, err
	}

	// Trace provider
	var traceProvider core.TraceProvider

	if cfg.Trace.Enabled {
		traceProvider = utils.NewTraceProvider()
	}

	// Stack trace provider
	var stackProvider core.StackTraceProvider

	if cfg.StackTrace.Enabled {
		stackProvider = utils.NewStackTraceProvider()
	}

	// Manager
	manager := core.NewManager(core.ManagerConfig{
		Logger:             logger,
		TraceProvider:      traceProvider,
		StackTraceProvider: stackProvider,
	})

	// DB Adapter
	var dbAdapter *pgxadapter.Adapter

	if cfg.Database.Type == "pgx" {

		dbAdapter = pgxadapter.New().
			WithConstraintDetails(cfg.Database.IncludeConstraintDetails).
			WithTableDetails(cfg.Database.IncludeTableDetails)
	}

	// Validator adapter
	var validatorAdapter *validatoradapter.Adapter

	if cfg.Validator.Enabled {
		validatorAdapter = validatoradapter.New()
	}

	// Framework initialization
	Init(manager, dbAdapter, validatorAdapter)

	return manager, nil
}
