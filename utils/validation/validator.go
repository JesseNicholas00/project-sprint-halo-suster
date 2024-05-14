package validation

import (
	"github.com/JesseNicholas00/HaloSuster/utils/validation/image"
	"github.com/JesseNicholas00/HaloSuster/utils/validation/phone"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type EchoValidator struct {
	validator *validator.Validate
}

func (e *EchoValidator) Validate(i interface{}) error {
	return e.validator.Struct(i)
}

var customFields = []customField{
	{
		Tag:       "phoneNumber",
		Validator: phone.ValidatePhoneNumber,
	},
	{
		Tag:       "imageExtension",
		Validator: image.ValidateImageExtension,
	},
}

type customField struct {
	Tag       string
	Validator validator.Func
}

func NewEchoValidator() echo.Validator {
	validator := validator.New(validator.WithRequiredStructEnabled())

	for _, customField := range customFields {
		validator.RegisterValidation(customField.Tag, customField.Validator)
	}

	return &EchoValidator{
		validator: validator,
	}
}
