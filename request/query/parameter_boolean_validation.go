package query

import (
	"fmt"
)

type ParameterBooleanValidator struct {
	Value *bool `json:"value"`
}

func (p ParameterBooleanValidator) Validate(value bool) error {
	if p.Value != nil && *p.Value != value {
		return fmt.Errorf("value [%v] does not equal %v", value, *p.Value)
	}
	return nil
}
