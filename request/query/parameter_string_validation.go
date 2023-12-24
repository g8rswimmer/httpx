package query

import (
	"fmt"
	"regexp"
)

const (
	RegExUUIDv4 = "^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"
)

type ParameterStringValidation struct {
	Value *string  `json:"value"`
	RegEx *string  `json:"regex"`
	OneOf []string `json:"one_of"`
}

func (p ParameterStringValidation) Validate(value string) error {
	if len(value) == 0 {
		return fmt.Errorf("value must be present")
	}
	if p.Value != nil && *p.Value != value {
		return fmt.Errorf("value [%s] does not equal %s", value, *p.Value)
	}
	if p.RegEx != nil {
		match, err := regexp.MatchString(*p.RegEx, value)
		switch {
		case err != nil:
			return fmt.Errorf("reg exp [%s] error %w", *p.RegEx, err)
		case !match:
			return fmt.Errorf("value [%s] does not match reg exp %s", value, *p.RegEx)
		default:
		}
	}
	if len(p.OneOf) == 0 {
		return nil
	}
	for _, s := range p.OneOf {
		if s == value {
			return nil
		}
	}
	return fmt.Errorf("value [%s] not in %v", value, p.OneOf)
}
