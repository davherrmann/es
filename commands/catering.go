package commands

import (
	"errors"
	"time"
)

// OrderFood command
type OrderFood struct {
	User  string
	Place string
	Date  time.Time
	Food  string
}

// Validate command
func (c *OrderFood) Validate() error {
	if c.Food == "" {
		return errors.New("Food must not be empty")
	}

	return nil
}
