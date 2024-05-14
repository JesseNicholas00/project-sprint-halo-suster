package nip

import (
	"github.com/JesseNicholas00/HaloSuster/types/nip"
	"github.com/go-playground/validator/v10"
)

func validateNipImpl(val int64) bool {
	return nip.IsValid(val)
}

func validateNipNurseImpl(val int64) bool {
	return nip.IsValid(val) && nip.GetRole(val) == nip.RoleNurse
}

func validateNipItImpl(val int64) bool {
	return nip.IsValid(val) && nip.GetRole(val) == nip.RoleIt
}

func ValidateNip(fl validator.FieldLevel) bool {
	return validateNipImpl(fl.Field().Int())
}

func ValidateNipNurse(fl validator.FieldLevel) bool {
	return validateNipNurseImpl(fl.Field().Int())
}

func ValidateNipIt(fl validator.FieldLevel) bool {
	return validateNipItImpl(fl.Field().Int())
}
