package image

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func validateImageExtenstion(imageUrl string) bool {
	for _, ext := range imageExtensions {
		if strings.HasSuffix(imageUrl, ext) {
			return true
		}
	}
	return false
}

func ValidateImageExtension(fl validator.FieldLevel) bool {
	return validateImageExtenstion(fl.Field().String())
}
