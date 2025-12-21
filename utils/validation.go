package utils

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) map[string]string {
	errorsMap := make(map[string]string)
	var ve validator.ValidationErrors

	if errors.As(err, &ve) {
		for _, fe := range ve {
			key := strings.ToLower(fe.Field())

			switch fe.Tag() {
			case "required":
				errorsMap[key] = "This field is required"
			case "email":
				errorsMap[key] = "Invalid email format"
			}
		}
	}
	return errorsMap
}