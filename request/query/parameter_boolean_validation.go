package query

import (
	"fmt"
)

type ParameterBooleanValidation struct {
	Value *bool `json:"value"`
}

func (p ParameterBooleanValidation) Validate(value bool) error {
	if p.Value != nil && *p.Value != value {
		return fmt.Errorf("value [%v] does not equal %v", value, *p.Value)
	}
	return nil
}
