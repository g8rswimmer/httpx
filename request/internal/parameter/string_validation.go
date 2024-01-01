package parameter

import (
	"errors"
	"fmt"
	"regexp"
)

const (
	RegExUUIDv4 = "^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$"
)

type StringValidator struct {
	Value *string  `json:"value"`
	RegEx *string  `json:"regex"`
	OneOf []string `json:"one_of"`
}

func (p StringValidator) Validate(value string) error {
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

type StringArrayValidator struct {
	Values  []string `json:"values"`
	RegEx   *string  `json:"regex"`
	Present []string `json:"present"`
}

func (s StringArrayValidator) Validate(values []string) error {
	if len(s.Values) > 0 {
		if len(s.Values) != len(values) {
			return errors.New("validator values lenght must match values length")
		}
		for i := range s.Values {
			if s.Values[i] != values[i] {
				return fmt.Errorf("value [%s] does not equal %s", values[i], s.Values[i])
			}
		}
	}
	if s.RegEx != nil {
		reqEx, err := regexp.Compile(*s.RegEx)
		if err != nil {
			return fmt.Errorf("reg exp [%s] error %w", *s.RegEx, err)
		}
		for _, value := range values {
			if !reqEx.MatchString(value) {
				return fmt.Errorf("value [%s] does not match reg exp %s", value, *s.RegEx)
			}
		}
	}
	if len(s.Present) == 0 {
		return nil
	}
	vset := map[string]struct{}{}
	for _, v := range values {
		vset[v] = struct{}{}
	}
	for _, p := range s.Present {
		if _, has := vset[p]; !has {
			return fmt.Errorf("value [%s] not in %v", p, values)
		}
	}

	return nil
}
