package query

import (
	"fmt"
	"strconv"

	"github.com/g8rswimmer/httpx/request/internal/parameter"
)

type BooleanValidator struct {
	parameter.BooleanValidator
}

func (p BooleanValidator) Validate(value string) error {
	b, err := strconv.ParseBool(value)
	if err != nil {
		return fmt.Errorf("query value is not a boolean [%s] %w", value, err)
	}
	return p.BooleanValidator.Validate(b)
}
