package datumcloud

import (
	"fmt"
)

var (
	// ErrDatumAPITokenMissing is returned when the Datum API token is missing
	ErrDatumAPITokenMissing = fmt.Errorf("DATUMCLOUD_TOKEN is required")
)

// RequiredFieldMissingError is returned when a field is required but not provided
type RequiredFieldMissingError struct {
	// Field contains the required field that was missing from the input
	Field string
}

// Error returns the RequiredFieldMissingError in string format
func (e *RequiredFieldMissingError) Error() string {
	return fmt.Sprintf("%s is required", e.Field)
}

// NewRequiredFieldMissingError returns an error for a missing required field
func NewRequiredFieldMissingError(f string) *RequiredFieldMissingError {
	return &RequiredFieldMissingError{
		Field: f,
	}
}
