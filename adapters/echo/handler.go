package echoadapter

import (
	"github.com/krisalay/error-framework/core"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	manager *core.Manager
}

func NewHandler(manager *core.Manager) *Handler {
	return &Handler{
		manager: manager,
	}
}

func (h *Handler) Handle(err error, c echo.Context) {

	ctx := c.Request().Context()

	appErr := h.manager.Handle(ctx, err)

	response := map[string]any{
		"message":  appErr.SafeMessage(),
		"code":     appErr.SafeCode(),
		"status":   appErr.Status,
		"trace_id": appErr.TraceID,
	}

	// include details only if not sensitive
	if !appErr.IsSensitive && len(appErr.Details) > 0 {
		response["details"] = appErr.Details
	}

	if !c.Response().Committed {
		c.JSON(appErr.Status, response)
	}
}
