package api

import (
	"time"
)

// TimeProvider returns a time
type TimeProvider func() time.Time

// DefaultTimeProvider returns the current UTC time
func DefaultTimeProvider() time.Time {
	return time.Now().UTC()
}
