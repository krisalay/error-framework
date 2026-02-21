package framework

import "github.com/krisalay/error-framework/core"

func Validation(err error) *core.AppError {

	if err == nil {
		return nil
	}

	f := get()

	return f.validatorAdapter.FromValidationError(err)
}
