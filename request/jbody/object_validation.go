package jbody

import (
	"errors"
	"fmt"

	"github.com/g8rswimmer/httpx/request/field"
	"github.com/g8rswimmer/httpx/request/rerror"
)

type validator interface {
	Validate(value any) error
}

type ObjectValidator struct {
	RequiredFields field.Required                 `json:"required_fields"`
	Parameters     map[string]ParameterProperties `json:"parameters"`
}

func (o ObjectValidator) Validate(value any) error {
	var obj map[string]any
	switch v := value.(type) {
	case map[string]any:
		obj = v
	default:
		return fmt.Errorf("value is not an object [%T]", value)
	}

	fields := field.Set(obj)

	if err := o.RequiredFields.Validate(fields); err != nil {
		return err
	}

	if err := field.Validate(fields, o.Parameters); err != nil {
		return err
	}

	for field, value := range obj {
		properties := o.Parameters[field]
		val, err := propertyValidator(properties.Validation)
		if err != nil {
			return &rerror.ParameterErr{
				Parameters: map[string]string{
					field: fmt.Sprintf("object validator: %v", err),
				},
			}
		}
		if err := val.Validate(value); err != nil {
			return &rerror.ParameterErr{
				Parameters: map[string]string{
					field: err.Error(),
				},
			}
		}
	}
	return nil
}

func propertyValidator(validation ParameterValidation) (validator, error) {
	var val validator
	if validation.String != nil {
		val = validation.String
	}
	if validation.StringArray != nil {
		if val != nil {
			return nil, errors.New("mulitple validators not allowed")
		}
		val = validation.StringArray
	}
	if validation.Number != nil {
		if val != nil {
			return nil, errors.New("mulitple validators not allowed")
		}
		val = validation.Number
	}
	if validation.NumberArray != nil {
		if val != nil {
			return nil, errors.New("mulitple validators not allowed")
		}
		val = validation.NumberArray
	}
	if validation.Time != nil {
		if val != nil {
			return nil, errors.New("mulitple validators not allowed")
		}
		val = validation.Time
	}
	if validation.TimeArray != nil {
		if val != nil {
			return nil, errors.New("mulitple validators not allowed")
		}
		val = validation.TimeArray
	}
	if validation.Boolean != nil {
		if val != nil {
			return nil, errors.New("mulitple validators not allowed")
		}
		val = validation.Boolean
	}
	if validation.Object != nil {
		if val != nil {
			return nil, errors.New("mulitple validators not allowed")
		}
		val = validation.Object
	}
	if validation.ObjectArray != nil {
		if val != nil {
			return nil, errors.New("mulitple validators not allowed")
		}
		val = validation.ObjectArray
	}
	if val == nil {
		return nil, errors.New("validator must be present")
	}
	return val, nil
}
