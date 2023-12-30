package jbody

import (
	"fmt"

	"github.com/g8rswimmer/httpx/request/parameter"
)

type NumberValidator struct {
	parameter.NumberValidator
}

func (n NumberValidator) Validate(value any) error {
	switch num := value.(type) {
	case float64:
		return n.NumberValidator.Validate(num)
	default:
		return fmt.Errorf("value is not a number [%T]", value)
	}
}

type NumberArrayValidator struct {
	parameter.NumberArrayValidator
}

func (n NumberArrayValidator) Validate(value any) error {
	var numArr []float64
	switch arr := value.(type) {
	case []any:
		for _, v := range arr {
			n, ok := v.(float64)
			if !ok {
				return fmt.Errorf("value is not a number [%T]", v)
			}
			numArr = append(numArr, n)
		}
	case []float64:
		numArr = arr
	default:
		return fmt.Errorf("value is not a number [%T]", value)
	}
	return n.NumberArrayValidator.Validate(numArr)
}
