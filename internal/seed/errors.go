package seed

import (
	"fmt"
)

var (
	// ErrDatumAPITokenMissing is returned when the Datum API token is missing
	ErrAPITokenMissing = fmt.Errorf("token is required but not provided")
)
