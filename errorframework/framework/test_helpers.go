package framework

import "github.com/krisalay/error-framework/core"

type mockLogger struct{}

func (m *mockLogger) Log(err *core.AppError) {}

func initTestFramework() {

	manager := core.NewManager(core.ManagerConfig{
		Logger: &mockLogger{},
	})

	Init(manager, nil, nil)
}
