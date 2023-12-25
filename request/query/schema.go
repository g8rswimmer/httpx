package query

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type SchemaError struct {
	Msg         string                         `json:"message"`
	MissingKeys []string                       `json:"missing_keys,omitempty"`
	Properties  map[string]SchemaPropertyError `json:"properties,omitempty"`
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
	Title       string                         `json:"title"`
	Description string                         `json:"description"`
	Parameters  map[string]ParameterProperties `json:"parameters"`
}

func (s Schema) Validate(req *http.Request) error {
	q, err := url.ParseQuery(req.URL.RawQuery)
	if err != nil {
		return &SchemaError{
			Msg: err.Error(),
		}
	}
	missingKeys := []string{}
	for k := range q {
		if _, has := s.Parameters[k]; !has {
			missingKeys = append(missingKeys, k)
		}
	}
	propertyErrors := map[string]SchemaPropertyError{}
	for key, properties := range s.Parameters {
		value := q.Get(key)
		if err := properties.Validate(value); err != nil {
			propertyErrors[key] = SchemaPropertyError{
				Msg: err.Error(),
			}
		}
	}
	if len(missingKeys) > 0 || len(propertyErrors) > 0 {
		return &SchemaError{
			Msg:         "schema validation error",
			MissingKeys: missingKeys,
			Properties:  propertyErrors,
		}
	}
	return nil
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
	for param, properties := range schema.Parameters {
		if err := SchemaModelParameterPropertiesValidator(properties); err != nil {
			return fmt.Errorf("schema parameter [%s]: %w", param, err)
		}
	}
	return nil
}
