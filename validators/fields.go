package validators

func ValidateAll(validationErrors ...ValidationErrors) ValidationErrors {
	var errors ValidationErrors

	for _, validationError := range validationErrors {
		errors = append(errors, validationError...)
	}
	return errors
}

func Validate(val interface{}, fieldName string, validators ...FieldValidator) ValidationErrors {
	var errors ValidationErrors

	for _, validator := range validators {
		errors = append(errors, validator(val, fieldName)...)
	}

	return errors
}
