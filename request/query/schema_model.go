package query

import (
	"errors"
	"fmt"
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
	Title           string                         `json:"title"`
	Description     string                         `json:"description"`
	LooseValidation bool                           `json:"loose_validation"`
	Parameters      map[string]ParameterProperties `json:"parameters"`
}

type ParameterProperties struct {
	Description          string `json:"description"`
	Example              string `json:"example"`
	DataType             string `json:"data_type"`
	InlineArray          bool   `json:"inline_array"`
	InlineArraySeperator string `json:"inline_array_seperator"`
}

func SchemaModelValidator(schema Schema) error {
	switch {
	case len(schema.Title) == 0:
		return errors.New("schema root title is required")
	case len(schema.Parameters) == 0:
		return errors.New("schema paremeters title is required")
	default:
	}
	for _, properties := range schema.Parameters {
		if err := SchemaModelParameterPropertiesValidator(properties); err != nil {
			return err
		}
	}
	return nil
}

func SchemaModelParameterPropertiesValidator(properties ParameterProperties) error {
	if _, has := dataTypes[properties.DataType]; !has {
		return fmt.Errorf("schema parameter propoerties data type error: %s not found", properties.DataType)
	}
	if properties.InlineArray && len(properties.InlineArraySeperator) == 0 {
		return errors.New("schema parameter properties inline array requires a seperator")
	}
	return nil
}
