package validation

import (
	"fmt"
	"time"
)

type ValidationError struct {
	Time    time.Time
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.Time, e.Message)
}
