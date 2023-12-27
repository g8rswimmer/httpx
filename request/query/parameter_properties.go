package query

import (
	"errors"
	"fmt"
	"strings"

	"github.com/g8rswimmer/httpx/request/parameter"
)

type ParameterValidation struct {
	String  *parameter.StringValidator `json:"string_validator"`
	Number  *NumberValidator           `json:"number_validator"`
	Time    *parameter.TimeValidator   `json:"time_validator"`
	Boolean *BooleanValidator          `json:"boolean_validator"`
}

func (p ParameterValidation) Validate(values []string) error {
	if err := p.validator(); err != nil {
		return err
	}
	for _, v := range values {
		switch {
		case p.String != nil:
			if err := p.String.Validate(v); err != nil {
				return err
			}
		case p.Number != nil:
			if err := p.Number.Validate(v); err != nil {
				return err
			}
		case p.Time != nil:
			if err := p.Time.Validate(v); err != nil {
				return err
			}
		case p.Boolean != nil:
			if err := p.Boolean.Validate(v); err != nil {
				return err
			}
		default:
			return errors.New("unable to validate the parameter")
		}
	}
	return nil
}

func (p ParameterValidation) validator() error {
	found := false
	if p.String != nil {
		found = true
	}
	if p.Number != nil {
		if found {
			return errors.New("parameter validation can't have more than one validator")
		}
		found = true
	}
	if p.Time != nil {
		if found {
			return errors.New("parameter validation can't have more than one validator")
		}
		found = true
	}
	if p.Boolean != nil {
		if found {
			return errors.New("parameter validation can't have more than one validator")
		}
		found = true
	}
	if !found {
		return errors.New("paramter validation must have one validation")
	}
	return nil
}

type ParameterProperties struct {
	Description          string              `json:"description"`
	Example              string              `json:"example"`
	InlineArray          bool                `json:"inline_array"`
	InlineArraySeperator string              `json:"inline_array_seperator"`
	Validation           ParameterValidation `json:"validation"`
}

func (p ParameterProperties) Validate(value string) error {
	var values []string
	switch {
	case len(value) == 0:
		return nil
	case p.InlineArray:
		values = strings.Split(value, p.InlineArraySeperator)
	default:
		values = []string{value}
	}

	if err := p.Validation.Validate(values); err != nil {
		return fmt.Errorf("query valiation: %w", err)
	}

	return nil
}

func SchemaModelParameterPropertiesValidator(properties ParameterProperties) error {
	if err := properties.Validation.validator(); err != nil {
		return fmt.Errorf("propoerties data type error: %w", err)
	}
	if properties.InlineArray && len(properties.InlineArraySeperator) == 0 {
		return errors.New("properties inline array requires a seperator")
	}
	return nil
}
