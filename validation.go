package monoacquiring

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var (
	// 01–12 і 00–99 mmyy
	cardExpRegex = regexp.MustCompile(`^(0[1-9]|1[0-2])[0-9]{2}$`)
)

func cardExpValidation(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	return cardExpRegex.MatchString(value)
}
