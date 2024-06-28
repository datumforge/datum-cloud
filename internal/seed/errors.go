package seed

import (
	"fmt"
)

var (
	// ErrDatumAPITokenMissing is returned when the Datum API token is missing
	ErrAPITokenMissing = fmt.Errorf("token is required but not provided")

	// ErrColumnNotFound is returned when a column is not found in the CSV file
	ErrColumnNotFound = fmt.Errorf("column not found in CSV file")

	// ErrInvalidTemplateName is returned when an invalid template name is provided
	ErrInvalidTemplateName = fmt.Errorf("invalid template name")
)
