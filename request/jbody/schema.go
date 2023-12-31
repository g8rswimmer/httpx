package jbody

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Schema struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Body        Body   `json:"body"`
}

func (s Schema) Validate(req *http.Request) error {
	var body any
	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		return fmt.Errorf("schema body json decode: %w", err)
	}
	if err := s.Body.Validate(body); err != nil {
		return fmt.Errorf("schema body err: %w", err)
	}
	return nil
}

func SchemaFromJSON(reader io.Reader) (Schema, error) {
	var schema Schema
	if err := json.NewDecoder(reader).Decode(&schema); err != nil {
		return Schema{}, fmt.Errorf("schema decode json: %w", err)
	}
	return schema, nil
}
