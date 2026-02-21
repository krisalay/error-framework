package main

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/krisalay/error-framework/errorframework/framework"
	"github.com/labstack/echo/v4"
)

func GetUserHandler(c echo.Context) error {

	err := errors.New("database timeout")

	return framework.WrapSafe(err, "failed to fetch user")
}

func GetPostHandler(c echo.Context) error {

	err := errors.New("database timeout")

	return framework.Wrap(err, "failed to fetch user")
}

type UserRequest struct {
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age" validate:"gte=18"`
}

func ValidationErrorHandler(c echo.Context) error {
	var validate = validator.New()
	req := UserRequest{
		Email: "invalid-email",
		Age:   10,
	}

	err := validate.Struct(req)

	if err != nil {
		return framework.Validation(err)
	}

	return nil
}
