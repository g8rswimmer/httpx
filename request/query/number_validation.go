package query

import (
	"fmt"
	"strconv"

	"github.com/g8rswimmer/httpx/request/parameter"
)

type NumberValidator struct {
	parameter.NumberValidator
}

func (p NumberValidator) Validate(value string) error {
	num, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return fmt.Errorf("query value is not a number [%s] %w", value, err)
	}
	return p.NumberValidator.Validate(num)
}
