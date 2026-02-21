package echoadapter

import (
	"net/http"

	"github.com/krisalay/error-framework/core"
	"github.com/labstack/echo/v4"
)

func PanicMiddleware(manager *core.Manager) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(c echo.Context) error {

			defer func() {

				if recovered := recover(); recovered != nil {

					ctx := c.Request().Context()

					appErr := manager.HandlePanic(ctx, recovered)

					response := map[string]any{
						"message":  appErr.SafeMessage(),
						"code":     appErr.SafeCode(),
						"status":   http.StatusInternalServerError,
						"trace_id": appErr.TraceID,
					}

					c.JSON(http.StatusInternalServerError, response)
				}

			}()

			return next(c)
		}
	}
}
