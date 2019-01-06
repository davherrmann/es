package command

import (
	"context"
	"errors"
)

// OrderFood command
type OrderFood struct {
	User  string
	Place string
	Date  string
	Food  string
}

// CancelFoodOrder command
type CancelFoodOrder struct {
	Place string
	Date  string
	User  string
}

// Validate command
func (c OrderFood) Validate(ctx context.Context) error {
	if c.Food == "" {
		return errors.New("food must not be empty")
	}

	return nil
}
