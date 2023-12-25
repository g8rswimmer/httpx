package query

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/g8rswimmer/httpx/request/field"
)

type SchemaError struct {
	Msg            string                         `json:"message"`
	MissingFields  []string                       `json:"missing_fields,omitempty"`
	RequiredFields string                         `json:"required_fields,omitempty"`
	Properties     map[string]SchemaPropertyError `json:"properties,omitempty"`
}

func (s SchemaError) Error() string {
	return s.Msg
}

func (s *SchemaError) Is(target error) bool {
	_, ok := target.(*SchemaError)
	return ok
}

type SchemaPropertyError struct {
	Msg string `json:"message"`
}

type Schema struct {
	Title          string                         `json:"title"`
	Description    string                         `json:"description"`
	RequiredFields field.Required                 `json:"required_fields"`
	Parameters     map[string]ParameterProperties `json:"parameters"`
}

func (s Schema) Validate(req *http.Request) error {
	values, err := url.ParseQuery(req.URL.RawQuery)
	if err != nil {
		return &SchemaError{
			Msg: err.Error(),
		}
	}
	fields := fieldsFromValues(values)
	missingFields := []string{}
	for field := range fields {
		if _, has := s.Parameters[field]; !has {
			missingFields = append(missingFields, field)
		}
	}

	requiredErr := s.RequiredFields.Validate(fields)

	propertyErrors := map[string]SchemaPropertyError{}
	for key, properties := range s.Parameters {
		value := values.Get(key)
		if err := properties.Validate(value); err != nil {
			propertyErrors[key] = SchemaPropertyError{
				Msg: err.Error(),
			}
		}
	}
	return handleValidationError(missingFields, propertyErrors, requiredErr)
}

func handleValidationError(missingFields []string, propertyErrors map[string]SchemaPropertyError, requiredErr error) error {
	if len(missingFields) == 0 && len(propertyErrors) == 0 && requiredErr == nil {
		return nil
	}
	schemErr := &SchemaError{
		Msg:           "schema validation error",
		MissingFields: missingFields,
		Properties:    propertyErrors,
		RequiredFields: func() string {
			if requiredErr == nil {
				return ""
			}
			return requiredErr.Error()
		}(),
	}
	return schemErr
}

func fieldsFromValues(values url.Values) map[string]struct{} {
	fields := map[string]struct{}{}
	for field := range values {
		fields[field] = struct{}{}
	}
	return fields
}
func SchemaFromJSON(reader io.Reader) (Schema, error) {
	var schema Schema
	if err := json.NewDecoder(reader).Decode(&schema); err != nil {
		return Schema{}, fmt.Errorf("schema decode json: %w", err)
	}
	if err := SchemaModelValidator(schema); err != nil {
		return Schema{}, fmt.Errorf("schema decode validation: %w", err)
	}
	return schema, nil
}

func SchemaModelValidator(schema Schema) error {
	switch {
	case len(schema.Parameters) == 0:
		return errors.New("schema parameters title is required")
	default:
	}
	parameters := map[string]struct{}{}
	for param, properties := range schema.Parameters {
		if err := SchemaModelParameterPropertiesValidator(properties); err != nil {
			return fmt.Errorf("schema parameter [%s]: %w", param, err)
		}
		parameters[param] = struct{}{}
	}
	if err := schema.RequiredFields.Validate(parameters); err != nil {
		return fmt.Errorf("scheam required paramters missing: %w", err)
	}
	return nil
}
