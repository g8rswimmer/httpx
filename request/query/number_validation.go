package query

import (
	"fmt"
	"strconv"

	"github.com/g8rswimmer/httpx/request/internal/parameter"
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

type NumberArrayValidator struct {
	parameter.NumberArrayValidator
}

func (n NumberArrayValidator) Validate(values []string) error {
	nums := make([]float64, len(values))
	for i, value := range values {
		num, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return fmt.Errorf("query value is not a number [%s] %w", value, err)
		}
		nums[i] = num
	}
	return n.NumberArrayValidator.Validate(nums)
}
