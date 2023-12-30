package body

import (
	"errors"
	"fmt"

	"github.com/g8rswimmer/httpx/request/field"
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

	fields := fieldsFromObject(obj)
	if err := o.RequiredFields.Validate(fields); err != nil {
		return fmt.Errorf("object validator required fields: %w", err)
	}

	extraFields := []string{}
	for field := range fields {
		if _, has := o.Parameters[field]; !has {
			extraFields = append(extraFields, field)
		}
	}
	if len(extraFields) > 0 {
		return fmt.Errorf("extra fields present in object %v", extraFields)
	}

	for field, value := range obj {
		properties := o.Parameters[field]
		val, err := propertyValidator(properties.Validation)
		if err != nil {
			return fmt.Errorf("object validator [%s]: %w", field, err)
		}
		if err := val.Validate(value); err != nil {
			return fmt.Errorf("object validation [%s]: %w", field, err)
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

func fieldsFromObject(obj map[string]any) map[string]struct{} {
	fields := map[string]struct{}{}
	for field := range obj {
		fields[field] = struct{}{}
	}
	return fields
}
