package query

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	DataTypeString  = "string"
	DataTypeNumber  = "number"
	DataTypeBoolean = "bool"
)

var dataTypes = map[string]struct{}{
	DataTypeString:  {},
	DataTypeNumber:  {},
	DataTypeBoolean: {},
}

type Schema struct {
	Title       string                         `json:"title"`
	Description string                         `json:"description"`
	Parameters  map[string]ParameterProperties `json:"parameters"`
}

func (s Schema) Validate(req *http.Request) error {
	q, err := url.ParseQuery(req.URL.RawQuery)
	if err != nil {
		return err
	}
	for k := range q {
		if _, has := s.Parameters[k]; !has {
			return errors.New("")
		}
	}
	for key, properties := range s.Parameters {
		value := q.Get(key)
		if err := properties.Validate(value); err != nil {
			return err
		}
	}
	return nil
}

type ParameterProperties struct {
	Description          string `json:"description"`
	Example              string `json:"example"`
	DataType             string `json:"data_type"`
	InlineArray          bool   `json:"inline_array"`
	InlineArraySeperator string `json:"inline_array_seperator"`
	Optional             bool   `json:"optional"`
}

func (p ParameterProperties) Validate(value string) error {
	var values []string
	switch {
	case p.Optional && len(value) == 0:
		return nil
	case len(value) == 0:
		return fmt.Errorf("query value must be present")
	case p.InlineArray:
		values = strings.Split(value, p.InlineArraySeperator)
	default:
		values = []string{value}
	}

	for _, v := range values {
		switch p.DataType {
		case DataTypeString:
		case DataTypeNumber:
			if _, err := strconv.ParseFloat(v, 64); err != nil {
				return fmt.Errorf("query value is not a number %s %w", v, err)
			}
		case DataTypeBoolean:
			if _, err := strconv.ParseBool(v); err != nil {
				return fmt.Errorf("query value is not a boolean %s %w", v, err)
			}
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
	case len(schema.Title) == 0:
		return errors.New("schema root title is required")
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

func SchemaModelParameterPropertiesValidator(properties ParameterProperties) error {
	if _, has := dataTypes[properties.DataType]; !has {
		return fmt.Errorf("propoerties data type error: %s not found", properties.DataType)
	}
	if properties.InlineArray && len(properties.InlineArraySeperator) == 0 {
		return errors.New("properties inline array requires a seperator")
	}
	return nil
}
