package jbody

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/g8rswimmer/httpx/request/rerror"
)

type Schema struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Body        Body   `json:"body"`
}

func (s Schema) Validate(req *http.Request) (any, error) {
	var body any
	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		return nil, rerror.SchemaFromError("request json body validation", fmt.Errorf("schema body json decode: %w", err))
	}
	if err := s.Body.Validate(body); err != nil {
		return nil, rerror.SchemaFromError("request json body validation", err)
	}
	return body, nil
}

func SchemaFromJSON(reader io.Reader) (Schema, error) {
	var schema Schema
	if err := json.NewDecoder(reader).Decode(&schema); err != nil {
		return Schema{}, fmt.Errorf("schema decode json: %w", err)
	}
	return schema, nil
}
