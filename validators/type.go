package validators

import "google.golang.org/genproto/googleapis/rpc/errdetails"

type ValidationError *errdetails.BadRequest_FieldViolation
type ValidationErrors []*errdetails.BadRequest_FieldViolation
type FieldValidator func(val interface{}, fieldName string) ValidationErrors

func NewValidationError(field, reason string) ValidationError {
	return &errdetails.BadRequest_FieldViolation{
		Field:       field,
		Description: reason,
	}
}
