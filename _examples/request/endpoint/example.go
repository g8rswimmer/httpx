package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/g8rswimmer/httpx/request/endpoint"
	"github.com/g8rswimmer/httpx/request/rerror"
)

func main() {
	f, err := os.Open("schema.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	schema, err := endpoint.SchemaFromJSON(f)
	if err != nil {
		panic(err)
	}

	RequestEndpointSuccess(schema)

	RequestQueryFail(schema)
}

func RequestEndpointSuccess(schema endpoint.Schema) {
	req := httptest.NewRequest(http.MethodGet, "https://www.query-example.com/schema/example/8cf69907-82f9-4504-8b72-9a608b6381ec", nil)

	if err := schema.Validate(req); err != nil {
		panic(err)
	}
	fmt.Println("HTTP request endpoint validated: success")
}

func RequestQueryFail(schema endpoint.Schema) {
	req := httptest.NewRequest(http.MethodGet, "https://www.query-example.com/schema/example/no-id", nil)

	err := schema.Validate(req)
	var schemaErr *rerror.SchemaErr
	switch {
	case err == nil:
		panic("error expected for the request query fail")
	case errors.As(err, &schemaErr):
		enc, _ := json.MarshalIndent(schemaErr, "", "  ")
		fmt.Println("HTTP request endpoint validated: failure")
		fmt.Println(string(enc))
	default:
		panic(err)
	}
}
