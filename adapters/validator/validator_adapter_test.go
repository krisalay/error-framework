package validatoradapter

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

type TestStruct struct {
	Email string `validate:"required,email"`
}

func TestValidationError(t *testing.T) {

	validate := validator.New()

	obj := TestStruct{}

	err := validate.Struct(obj)

	adapter := New()

	appErr := adapter.FromValidationError(err)

	if appErr.Code == "" {
		t.Fatal("Validation adapter failed")
	}

	if len(appErr.Details) == 0 {
		t.Fatal("Details missing")
	}
}
