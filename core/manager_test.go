package core

import (
	"errors"
	"testing"
)

type mockLogger struct {
	called bool
}

func (m *mockLogger) Log(err *AppError) {
	m.called = true
}
func TestManager_Handle_AppError(t *testing.T) {

	logger := &mockLogger{}

	manager := NewManager(ManagerConfig{
		Logger: logger,
	})

	err := New().
		WithMessage("test").
		Build()

	result := manager.Handle(nil, err)

	if result.Message != "test" {
		t.Fatal("Handle failed")
	}

	if !logger.called {
		t.Fatal("Logger not called")
	}
}

func TestManager_Handle_UnknownError(t *testing.T) {

	logger := &mockLogger{}

	manager := NewManager(ManagerConfig{
		Logger: logger,
	})

	result := manager.Handle(nil, errors.New("unknown"))

	if result.Code != CodeInternalError {
		t.Fatal("Did not convert to internal error")
	}
}
