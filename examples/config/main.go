package main

import (
	"context"
	"errors"

	"github.com/krisalay/error-framework/core"
	"github.com/krisalay/error-framework/errorframework/framework"
	"github.com/krisalay/error-framework/logging"
)

func main() {

	// configure logger
	logger, err := logging.NewZapLogger(logging.Config{
		ConsoleEnabled: true,
		FileEnabled:    false,
		Level:          "debug",
	})

	if err != nil {
		panic(err)
	}

	// configure manager
	manager := core.NewManager(core.ManagerConfig{
		Logger: logger,
	})

	// initialize framework
	framework.Init(manager, nil, nil)

	// simulate error
	err = errors.New("database connection failed")

	appErr := framework.Wrap(err, "failed to fetch user")

	manager.Handle(context.TODO(), appErr)
}
