package main

import (
	"errors"
	"fmt"

	"github.com/krisalay/error-framework/errorframework/framework"
)

func main() {

	err := errors.New("connection refused")

	appErr := framework.Wrap(err, "failed to fetch user")

	fmt.Println(appErr.SafeMessage())
}
