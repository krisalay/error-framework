package framework

import "github.com/krisalay/error-framework/core"

func DB(err error) *core.AppError {

	if err == nil {
		return nil
	}

	f := get()

	return f.dbAdapter.FromError(err)
}
