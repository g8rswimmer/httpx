package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/g8rswimmer/httpx/request/endpoint"
	"github.com/g8rswimmer/httpx/request/jbody"
	"github.com/g8rswimmer/httpx/request/query"
	"github.com/g8rswimmer/httpx/request/rerror"
)

type SchemaResult struct {
	Path   string
	Values url.Values
	Body   any
}

type Schema struct {
	Endpoint *endpoint.Schema `json:"endpoint"`
	Query    *query.Schema    `json:"query"`
	Body     *jbody.Schema    `json:"body"`
}

func (s Schema) Validate(req *http.Request) (*SchemaResult, error) {
	path, err := s.validateEndpoint(req)
	if err != nil {
		return nil, err
	}
	values, err := s.validateQuery(req)
	if err != nil {
		return nil, err
	}
	body, err := s.validateBody(req)
	if err != nil {
		return nil, err
	}
	return &SchemaResult{
		Path:   path,
		Values: values,
		Body:   body,
	}, nil
}

func (s Schema) validateEndpoint(req *http.Request) (string, error) {
	switch {
	case s.Endpoint != nil:
		if err := s.Endpoint.Validate(req); err != nil {
			return "", err
		}
	default:
	}
	return req.URL.Path, nil
}

func (s Schema) validateQuery(req *http.Request) (url.Values, error) {
	var (
		values url.Values
		err    error
	)

	switch {
	case s.Query != nil:
		values, err = s.Query.Validate(req)
		if err != nil {
			return nil, err
		}
	default:
		values, err = url.ParseQuery(req.URL.RawQuery)
		if err != nil {
			return nil, rerror.SchemaFromError("request query parse", err)
		}
	}
	return values, nil
}

func (s Schema) validateBody(req *http.Request) (any, error) {
	var (
		body any
		err  error
	)
	switch {
	case s.Body != nil:
		body, err = s.Body.Validate(req)
		if err != nil {
			return nil, err
		}
	default:
		err = json.NewDecoder(req.Body).Decode(&body)
		switch {
		case errors.Is(err, io.EOF):
		case err != nil:
			return nil, rerror.SchemaFromError("request body decode", err)
		default:
		}
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
