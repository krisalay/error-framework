package main

import (
	"errors"

	"github.com/krisalay/error-framework/core"
	"github.com/krisalay/error-framework/logging"
)

func main() {

	logger, _ := logging.NewZapLogger(logging.Config{
		ConsoleEnabled: true,
		FileEnabled:    true,
		FilePath:       "error.log",
	})

	manager := core.NewManager(core.ManagerConfig{
		Logger: logger,
	})

	manager.Handle(nil, errors.New("test error"))
}
