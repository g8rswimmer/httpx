package endpoint

import (
	"errors"
	"fmt"
)

type PathVariable struct {
	Validation VariableValidation `json:"validation"`
}

func (pv PathVariable) Validate(value string) error {
	return pv.Validation.Validate(value)
}

type VariableValidation struct {
	String *StringValidator `json:"string_validator"`
	Number *NumberValidator `json:"number_validator"`
}

func (v VariableValidation) Validate(value string) error {
	if err := v.validator(); err != nil {
		return err
	}
	switch {
	case v.String != nil:
		if err := v.String.Validate(value); err != nil {
			return err
		}
	case v.Number != nil:
		if err := v.Number.Validate(value); err != nil {
			return err
		}
	default:
		return errors.New("unable to validate the parameter")
	}
	return nil
}

func (v VariableValidation) validator() error {
	found := false
	if v.String != nil {
		found = true
	}
	if v.Number != nil {
		if found {
			return errors.New("path validation can't have more than one validator")
		}
		found = true
	}
	if !found {
		return errors.New("path validation must have one validation")
	}
	return nil
}

func SchemaModelPathVariableValidator(pathVariable PathVariable) error {
	if err := pathVariable.Validation.validator(); err != nil {
		return fmt.Errorf("propoerties data type error: %w", err)
	}
	return nil
}
