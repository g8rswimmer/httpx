package query

import (
	"fmt"
	"strconv"

	"github.com/g8rswimmer/httpx/request/parameter"
)

type QueryBooleanValidator struct {
	parameter.BooleanValidator
}

func (p QueryBooleanValidator) Validate(value string) error {
	b, err := strconv.ParseBool(value)
	if err != nil {
		return fmt.Errorf("query value is not a boolean [%s] %w", value, err)
	}
	return p.BooleanValidator.Validate(b)
}
