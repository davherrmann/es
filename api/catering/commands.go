package catering

import (
	"context"
	"errors"
)

// DoOrderFood command
type DoOrderFood struct {
	User  string
	Place string
	Date  string
	Food  string
}

// DoCancelFoodOrder command
type DoCancelFoodOrder struct {
	Place string
	Date  string
	User  string
}

// Validate command
func (c DoOrderFood) Validate(ctx context.Context) error {
	if c.Food == "" {
		return errors.New("food must not be empty")
	}

	return nil
}
