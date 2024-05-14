package validation

import (
	"github.com/JesseNicholas00/HaloSuster/utils/validation/nip"
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
		Tag:       "phoneNum",
		Validator: phone.ValidatePhoneNumber,
	},
	{
		Tag:       "nip",
		Validator: nip.ValidateNip,
	},
	{
		Tag:       "nipNurse",
		Validator: nip.ValidateNipNurse,
	},
	{
		Tag:       "nipIt",
		Validator: nip.ValidateNipIt,
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
