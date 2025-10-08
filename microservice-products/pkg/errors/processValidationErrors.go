package errors

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// ProcessValidationErrors processes validator.ValidationErrors as errors.ValidationError.
func ProcessValidationErrors(ve validator.ValidationErrors) []ValidationError {
	var validationErrors []ValidationError

	for _, fe := range ve {
		validationError := &ValidationErrorImpl{
			field: strings.ToLower(fe.Field()),
			err:   getValidationErrorMessage(fe),
		}
		validationErrors = append(validationErrors, validationError)
	}

	return validationErrors
}

// getValidationErrorMessage returns a user-friendly error message based on the validation tag.
func getValidationErrorMessage(fe validator.FieldError) error {
	switch fe.Tag() {
	case "required":
		return fmt.Errorf("это поле обязательно для заполнения")
	case "max":
		return fmt.Errorf("максимальная длина %s символов", fe.Param())
	case "min":
		return fmt.Errorf("минимальная длина %s символов", fe.Param())
	case "gt":
		return fmt.Errorf("значение должно быть больше %s", fe.Param())
	case "gte":
		return fmt.Errorf("значение должно быть больше или равно %s", fe.Param())
	case "lt":
		return fmt.Errorf("значение должно быть меньше %s", fe.Param())
	case "lte":
		return fmt.Errorf("значение должно быть меньше или равно %s", fe.Param())
	case "email":
		return fmt.Errorf("неверный формат email")
	default:
		return fmt.Errorf("недопустимое значение")
	}
}
