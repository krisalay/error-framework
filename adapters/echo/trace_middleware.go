package echoadapter

import (
	"github.com/google/uuid"
	"github.com/krisalay/error-framework/utils"
	"github.com/labstack/echo/v4"
)

const (
	HeaderTraceID     = "X-Trace-ID"
	HeaderRequestID   = "X-Request-ID"
	HeaderTraceParent = "traceparent"
)

func TraceMiddleware(traceProvider *utils.TraceProvider) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(c echo.Context) error {

			req := c.Request()

			traceID := extractTraceID(c)

			ctx := traceProvider.Inject(req.Context(), traceID)

			c.SetRequest(req.WithContext(ctx))

			c.Response().Header().Set(HeaderTraceID, traceID)

			return next(c)
		}
	}
}

func extractTraceID(c echo.Context) string {

	traceID := c.Request().Header.Get(HeaderTraceID)

	if traceID != "" {
		return traceID
	}

	traceID = c.Request().Header.Get(HeaderRequestID)

	if traceID != "" {
		return traceID
	}

	traceID = c.Request().Header.Get(HeaderTraceParent)

	if traceID != "" {
		return traceID
	}

	return uuid.New().String()
}
