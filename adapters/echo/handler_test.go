package echoadapter

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/krisalay/error-framework/core"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// mockLogger implements core.Logger for testing
type mockLogger struct{}

func (m *mockLogger) Log(err *core.AppError) {}

// mockTraceProvider implements core.TraceProvider for testing
type mockTraceProvider struct{}

func (m *mockTraceProvider) GetTraceID(ctx context.Context) string {
	return "test-trace-id"
}

func TestHandler_Handle(t *testing.T) {
	// Create manager and handler with mock logger
	manager := core.NewManager(core.ManagerConfig{
		Logger:        &mockLogger{},
		TraceProvider: &mockTraceProvider{},
	})
	handler := NewHandler(manager)

	// Test case 1: Handle a normal error
	t.Run("handle normal error", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := &core.AppError{
			Status: http.StatusInternalServerError,
			Code:   "TEST_CODE",
			Message: "test error message",
		}
		handler.Handle(err, c)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Contains(t, rec.Body.String(), "test error message")
		assert.Contains(t, rec.Body.String(), "TEST_CODE")
	})

	// Test case 2: Handle error with details (non-sensitive)
	t.Run("handle error with details", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := &core.AppError{
			Status:  http.StatusBadRequest,
			Code:    "INVALID_INPUT",
			Message: "validation failed",
			Details: map[string]interface{}{"field": "username"},
		}
		handler.Handle(err, c)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Contains(t, rec.Body.String(), "username")
	})

	// Test case 3: Handle error with sensitive flag (details should be hidden)
	t.Run("handle sensitive error without details", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := &core.AppError{
			Status:     http.StatusInternalServerError,
			Code:       "INTERNAL_ERROR",
			Message:    "internal server error",
			IsSensitive: true,
			Details:    map[string]interface{}{"secret": "password123"},
		}
		handler.Handle(err, c)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.NotContains(t, rec.Body.String(), "password123")
	})

	// Test case 4: Response already committed (should not panic)
	t.Run("handle when response already committed", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Commit the response first
		rec.WriteHeader(http.StatusOK)

		err := &core.AppError{
			Status:  http.StatusInternalServerError,
			Code:    "TEST_CODE",
			Message: "test error",
		}
		
		// Should not panic
		handler.Handle(err, c)
	})
}
