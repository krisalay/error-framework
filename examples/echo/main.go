package main

import (
	echoadapter "github.com/krisalay/error-framework/adapters/echo"
	"github.com/krisalay/error-framework/errorframework/config"
	"github.com/krisalay/error-framework/errorframework/framework"
	"github.com/krisalay/error-framework/utils"
	"github.com/labstack/echo/v4"
)

func main() {

	// Initialize framework using config
	cfg := config.Config{
		Logger: config.LoggerConfig{
			ConsoleEnabled: true,
			FileEnabled:    false,
			Level:          "debug",
			Encoding:       "console",
		},
		Trace: config.TraceConfig{
			Enabled: true,
		},
		StackTrace: config.StackTraceConfig{
			Enabled: true,
		},
		Database: config.DatabaseConfig{
			Type:                     "pgx",
			IncludeTableDetails:      true,
			IncludeConstraintDetails: true,
		},
		Validator: config.ValidatorConfig{
			Enabled: true,
		},
	}

	manager, err := framework.InitFromConfig(cfg)
	if err != nil {
		panic(err)
	}

	traceProvider := utils.NewTraceProvider()

	// Echo setup
	e := echo.New()

	e.Use(echoadapter.TraceMiddleware(traceProvider))
	e.Use(echoadapter.PanicMiddleware(manager))

	e.HTTPErrorHandler = echoadapter.NewHandler(manager).Handle

	e.GET("/user", GetUserHandler)
	e.GET("/post", GetPostHandler)
	e.GET("/validation-error", ValidationErrorHandler)

	e.Logger.Info("Demo server running on :8080")
	e.Start(":8080")
}
