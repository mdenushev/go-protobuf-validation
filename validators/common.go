package validators

import "fmt"

func AllowedValues(values ...interface{}) FieldValidator {
	return func(val interface{}, fieldName string) ValidationErrors {
		var errors ValidationErrors
		if !IsInArray(val, values) {
			errors = append(errors, NewValidationError(fieldName, fmt.Sprintf("any:allowed")))
		}
		return errors
	}
}

func Required() FieldValidator {
	return func(val interface{}, fieldName string) ValidationErrors {
		var errors ValidationErrors

		if val == nil {
			errors = append(errors, NewValidationError(fieldName, fmt.Sprintf("any:required")))
		}
		return errors
	}
}
