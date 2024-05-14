package phone

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

var countryCodesByLen = func() []map[string]struct{} {
	res := make([]map[string]struct{}, maxCountryCodeLen+1)

	for i := range res {
		res[i] = make(map[string]struct{})
	}

	for _, code := range countryCodes {
		res[len(code)][code] = struct{}{}
	}

	return res
}()

func validatePhoneNumberImpl(phoneNum string) bool {
	// length check
	if len(phoneNum) < 10 || len(phoneNum) > 16 {
		return false
	}

	// should start with +
	if phoneNum[0] != '+' {
		return false
	}

	// has stray characters
	for _, c := range phoneNum[1:] {
		if !unicode.IsDigit(c) && c != '-' {
			return false
		}
	}

	// iterate by length
	for curLen, countryCodeSet := range countryCodesByLen {
		if _, ok := countryCodeSet[phoneNum[1:curLen+1]]; ok {
			// found code in set
			return true
		}
	}

	return false
}

func ValidatePhoneNumber(fl validator.FieldLevel) bool {
	return validatePhoneNumberImpl(fl.Field().String())
}
