package validatoradapter

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/krisalay/error-framework/core"
)

type Adapter struct {
	messages map[string]string
}

func New() *Adapter {
	return &Adapter{
		messages: make(map[string]string),
	}
}

func (a *Adapter) RegisterMessage(tag string, message string) {
	a.messages[tag] = message
}

func (a *Adapter) RegisterFieldMessage(field string, tag string, message string) {
	key := field + "." + tag
	a.messages[key] = message
}

// FromValidationError converts validator.ValidationErrors to AppError
func (a *Adapter) FromValidationError(err error) *core.AppError {

	validationErrors, ok := err.(validator.ValidationErrors)

	if !ok {
		return core.New().
			WithMessage("Validation failed").
			WithCode(core.CodeValidationError).
			WithStatus(http.StatusBadRequest).
			WithSensitive(false).
			Build()
	}

	details := make(map[string]any)

	for _, fieldErr := range validationErrors {

		field := toSnakeCase(fieldErr.Field())

		message := a.getMessage(fieldErr)

		details[field] = message
	}

	return core.New().
		WithMessage("Validation failed").
		WithCode(core.CodeValidationError).
		WithStatus(http.StatusBadRequest).
		WithDetails(details).
		WithLevel(core.LevelWarn).
		WithSensitive(false).
		Build()
}

func (a *Adapter) getMessage(err validator.FieldError) string {

	field := toSnakeCase(err.Field())
	tag := err.Tag()

	// Check field-specific message first
	fieldKey := field + "." + tag

	if msg, ok := a.messages[fieldKey]; ok {
		return msg
	}

	// Check tag-level message
	if msg, ok := a.messages[tag]; ok {
		return msg
	}

	// Default fallback messages
	switch tag {

	case "required":
		return "is required"

	case "email":
		return "must be a valid email"

	case "gte":
		return fmt.Sprintf("must be >= %s", err.Param())

	case "lte":
		return fmt.Sprintf("must be <= %s", err.Param())

	case "min":
		return fmt.Sprintf("must be at least %s characters", err.Param())

	case "max":
		return fmt.Sprintf("must be at most %s characters", err.Param())

	default:
		return fmt.Sprintf("failed validation: %s", tag)
	}
}

func toSnakeCase(str string) string {

	var result []rune

	for i, r := range str {

		if i > 0 && r >= 'A' && r <= 'Z' {
			result = append(result, '_')
		}

		result = append(result, r)
	}

	return strings.ToLower(string(result))
}
