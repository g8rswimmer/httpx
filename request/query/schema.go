package query

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/g8rswimmer/httpx/request/field"
	"github.com/g8rswimmer/httpx/request/rerror"
)

type Schema struct {
	Title          string                         `json:"title"`
	Description    string                         `json:"description"`
	RequiredFields field.Required                 `json:"required_fields"`
	Parameters     map[string]ParameterProperties `json:"parameters"`
}

func (s Schema) Validate(req *http.Request) error {
	values, err := url.ParseQuery(req.URL.RawQuery)
	if err != nil {
		return rerror.SchemaFromError("request query validation", err)
	}

	if err := field.Validate(values, s.Parameters); err != nil {
		return rerror.SchemaFromError("request query validation", err)
	}

	set := field.Set(values)

	if err := s.RequiredFields.Validate(set); err != nil {
		return rerror.SchemaFromError("request query validation", err)
	}

	parameterErr := &rerror.ParameterErr{
		Parameters: map[string]string{},
	}
	for key, properties := range s.Parameters {
		value := values.Get(key)
		if err := properties.Validate(value); err != nil {
			parameterErr.Add(key, err.Error())
		}
	}
	return rerror.SchemaFromError("request query validation", parameterErr)
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
