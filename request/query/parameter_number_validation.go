package query

import (
	"fmt"
)

type ParameterDataNumberValidation struct {
	Value *float64
	Min   *float64  `json:"min"`
	Max   *float64  `json:"max"`
	OneOf []float64 `json:"one_of"`
}

func (p ParameterDataNumberValidation) Validate(num float64) error {
	if p.Value != nil && num != *p.Value {
		return fmt.Errorf("query value [%f] does not equal %f", num, *p.Value)
	}
	if p.Min != nil && num < *p.Min {
		return fmt.Errorf("query value [%f] is less than %f", num, *p.Min)
	}
	if p.Max != nil && num > *p.Max {
		return fmt.Errorf("query value [%f] is greater than %f", num, *p.Max)
	}
	if len(p.OneOf) == 0 {
		return nil
	}
	for _, n := range p.OneOf {
		if num == n {
			return nil
		}
	}
	return fmt.Errorf("query value [%f] not in %v", num, p.OneOf)
}