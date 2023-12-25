package parameter

import (
	"fmt"
)

type NumberValidator struct {
	Value *float64  `json:"value"`
	Min   *float64  `json:"min"`
	Max   *float64  `json:"max"`
	OneOf []float64 `json:"one_of"`
}

func (p NumberValidator) Validate(num float64) error {
	if p.Value != nil && num != *p.Value {
		return fmt.Errorf("value [%f] does not equal %f", num, *p.Value)
	}
	if p.Min != nil && num < *p.Min {
		return fmt.Errorf("value [%f] is less than %f", num, *p.Min)
	}
	if p.Max != nil && num > *p.Max {
		return fmt.Errorf("value [%f] is greater than %f", num, *p.Max)
	}
	if len(p.OneOf) == 0 {
		return nil
	}
	for _, n := range p.OneOf {
		if num == n {
			return nil
		}
	}
	return fmt.Errorf("value [%f] not in %v", num, p.OneOf)
}
