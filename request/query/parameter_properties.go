package query

import (
	"errors"
	"fmt"
	"strings"

	"github.com/g8rswimmer/httpx/request/internal/parameter"
)

type ParameterValidation struct {
	String      *parameter.StringValidator      `json:"string_validator"`
	Number      *NumberValidator                `json:"number_validator"`
	Time        *parameter.TimeValidator        `json:"time_validator"`
	Boolean     *BooleanValidator               `json:"boolean_validator"`
	StringArray *parameter.StringArrayValidator `json:"string_array_validator"`
	TimeArray   *parameter.TimeArrayValidator   `json:"time_array_validator"`
	NumberArray *NumberArrayValidator           `json:"number_array_validator"`
}

func (p ParameterValidation) ValidateValues(values []string) error {
	if err := p.validatorValues(); err != nil {
		return err
	}
	switch {
	case p.StringArray != nil:
		if err := p.StringArray.Validate(values); err != nil {
			return err
		}
	case p.NumberArray != nil:
		if err := p.NumberArray.Validate(values); err != nil {
			return err
		}
	case p.TimeArray != nil:
		if err := p.TimeArray.Validate(values); err != nil {
			return err
		}
	default:
		return errors.New("unable to validate the parameter")
	}
	return nil
}

func (p ParameterValidation) ValidateValue(value string) error {
	if err := p.validatorValue(); err != nil {
		return err
	}
	switch {
	case p.String != nil:
		if err := p.String.Validate(value); err != nil {
			return err
		}
	case p.Number != nil:
		if err := p.Number.Validate(value); err != nil {
			return err
		}
	case p.Time != nil:
		if err := p.Time.Validate(value); err != nil {
			return err
		}
	case p.Boolean != nil:
		if err := p.Boolean.Validate(value); err != nil {
			return err
		}
	default:
		return errors.New("unable to validate the parameter")
	}
	return nil
}

func (p ParameterValidation) validatorValue() error {
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

func (p ParameterValidation) validatorValues() error {
	found := false
	if p.StringArray != nil {
		found = true
	}
	if p.NumberArray != nil {
		if found {
			return errors.New("parameter validation can't have more than one validator")
		}
		found = true
	}
	if p.TimeArray != nil {
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
	if p.StringArray != nil {
		if found {
			return errors.New("parameter validation can't have more than one validator")
		}
		found = true
	}
	if p.NumberArray != nil {
		if found {
			return errors.New("parameter validation can't have more than one validator")
		}
		found = true
	}
	if p.TimeArray != nil {
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
	switch {
	case len(value) == 0:
		return nil
	case p.InlineArray:
		values := strings.Split(value, p.InlineArraySeperator)
		if err := p.Validation.ValidateValues(values); err != nil {
			return fmt.Errorf("query valiation: %w", err)
		}
	default:
		if err := p.Validation.ValidateValue(value); err != nil {
			return fmt.Errorf("query valiation: %w", err)
		}
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
