package main

import (
	"errors"

	"github.com/go-playground/validator/v10"
	echoadapter "github.com/krisalay/error-framework/adapters/echo"
	validatoradapter "github.com/krisalay/error-framework/adapters/validator"
	"github.com/krisalay/error-framework/errorframework/config"
	"github.com/krisalay/error-framework/errorframework/framework"
	"github.com/krisalay/error-framework/utils"
	"github.com/labstack/echo/v4"
)

var validate = validator.New()
var validatorAdapter = validatoradapter.New()

// Request struct for validation demo
type UserRequest struct {
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age" validate:"gte=18"`
}

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
			Type: "pgx",
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

	registerRoutes(e)

	e.Logger.Info("Demo server running on :8080")
	e.Start(":8080")
}

func registerRoutes(e *echo.Echo) {

	e.GET("/success", successHandler)

	e.GET("/validation-error", validationErrorHandler)

	e.GET("/db-error", dbErrorHandler)

	e.GET("/wrapped-error", wrappedErrorHandler)

	e.GET("/internal-error", internalErrorHandler)

	e.GET("/panic", panicHandler)

	e.GET("/custom-error", customErrorHandler)

	// e.GET("/goroutine-panic", goroutinePanicHandler)
}

// SUCCESS DEMO
func successHandler(c echo.Context) error {

	return c.JSON(200, map[string]string{
		"message": "success",
	})
}

// VALIDATION ERROR DEMO
func validationErrorHandler(c echo.Context) error {

	req := UserRequest{
		Email: "invalid-email",
		Age:   10,
	}

	err := validate.Struct(req)

	if err != nil {
		return framework.Validation(err)
	}

	return nil
}

// DATABASE ERROR DEMO (SIMULATED)
func dbErrorHandler(c echo.Context) error {

	dbErr := errors.New("duplicate key value violates unique constraint")

	return framework.DB(dbErr)
}

// WRAPPED ERROR DEMO
func wrappedErrorHandler(c echo.Context) error {

	originalErr := errors.New("connection timeout")

	return framework.Wrap(originalErr, "failed to fetch user")
}

// INTERNAL ERROR DEMO
func internalErrorHandler(c echo.Context) error {

	err := errors.New("something unexpected happened")

	return framework.Internal(err)
}

// PANIC DEMO
func panicHandler(c echo.Context) error {

	panic("simulated panic")

}

// CUSTOM ERROR DEMO
func customErrorHandler(c echo.Context) error {

	return framework.NotFound("user not found")
}

// GOROUTINE PANIC DEMO
// func goroutinePanicHandler(c echo.Context) error {

// 	ctx := c.Request().Context()

// 	go func() {

// 		defer framework.Recover(ctx)

// 		time.Sleep(1 * time.Second)

// 		panic("goroutine panic demo")

// 	}()

// 	return c.JSON(http.StatusOK, map[string]string{
// 		"message": "goroutine started",
// 	})
// }
