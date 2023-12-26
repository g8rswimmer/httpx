package endpoint

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Schema struct {
	Title         string                  `json:"title"`
	Description   string                  `json:"description"`
	Method        string                  `json:"method"`
	Endpoint      string                  `json:"endpoint"`
	PathVariables map[string]PathVariable `json:"path_variables"`
}

func (s Schema) Validate(req *http.Request) error {
	if req.Method != s.Method {
		return fmt.Errorf("request method [%s] does not match expected method [%s]", req.Method, s.Method)
	}
	reqPaths := strings.Split(req.URL.Path, "/")
	schemaPaths := strings.Split(s.Endpoint, "/")
	if len(reqPaths) != len(schemaPaths) {
		return fmt.Errorf("request paths do not match")
	}
	for i := range schemaPaths {
		if pv, has := s.PathVariables[schemaPaths[i]]; has {
			if err := pv.Validate(reqPaths[i]); err != nil {
				return err
			}
			continue
		}
		if schemaPaths[i] != reqPaths[i] {
			return fmt.Errorf("request path [%s] does not match [%s]", reqPaths[i], schemaPaths[i])
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
	case len(schema.Method) == 0:
		return errors.New("schema endpoint method is required")
	default:
	}
	for param, pathVariable := range schema.PathVariables {
		if err := SchemaModelPathVariableValidator(pathVariable); err != nil {
			return fmt.Errorf("schema parameter [%s]: %w", param, err)
		}
	}
	return nil
}