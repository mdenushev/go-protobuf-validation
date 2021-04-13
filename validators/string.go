package validators

import (
	"fmt"
	"regexp"
)

const (
	uuidRegexp  = "[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}"
	emailRegexp = ".+@.+\\..+"
)

func StringLength(min, max int) FieldValidator {
	return func(val interface{}, fieldName string) ValidationErrors {
		var errors ValidationErrors

		if len(val.(string)) < min {
			errors = append(errors, NewValidationError(fieldName, fmt.Sprintf("str:len_min:%d", min)))
		}

		if len(val.(string)) > max && max != 0 {
			errors = append(errors, NewValidationError(fieldName, fmt.Sprintf("str:len_max:%d", max)))
		}

		return errors
	}
}

func StringRegexp(regexp *regexp.Regexp, customError ...string) FieldValidator {
	return func(val interface{}, fieldName string) ValidationErrors {
		var errors ValidationErrors
		if !regexp.Match([]byte(val.(string))) {
			if len(customError) != 0 {
				errors = append(errors, NewValidationError(fieldName, fmt.Sprintf("regexp:%s", customError[0])))
			} else {
				errors = append(errors, NewValidationError(fieldName, fmt.Sprintf("regexp:%s", regexp.String())))
			}
		}
		return errors
	}
}

func UUIDRegexp() FieldValidator {
	return StringRegexp(regexp.MustCompile(uuidRegexp), "uuid")
}

func EmailRegexp() FieldValidator {
	return StringRegexp(regexp.MustCompile(emailRegexp), "email")
}
