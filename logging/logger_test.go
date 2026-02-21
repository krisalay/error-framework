package logging

import (
	"testing"

	"github.com/krisalay/error-framework/core"
)

func TestZapLogger(t *testing.T) {

	logger, err := NewZapLogger(Config{
		ConsoleEnabled: true,
		Level:          "debug",
	})

	if err != nil {
		t.Fatal(err)
	}

	logger.Log(core.New().WithMessage("test").Build())
}
