package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ConvertToErrList(err validator.ValidationErrors) (errors []string) {
	for _, err := range err {
		curError := fmt.Sprintf(
			"%s: validation failed on %s %s",
			err.Field(),
			err.Tag(),
			err.Param(),
		)
		errors = append(errors, curError)
	}
	return
}
