package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/g8rswimmer/httpx/request/query"
	"github.com/g8rswimmer/httpx/request/rerror"
)

func main() {
	f, err := os.Open("schema.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	schema, err := query.SchemaFromJSON(f)
	if err != nil {
		panic(err)
	}

	RequestQuerySuccess(schema)

	RequestQueryFail(schema)
}

func RequestQuerySuccess(schema query.Schema) {
	req := httptest.NewRequest(http.MethodGet, "https://www.query-example.com", nil)
	q := req.URL.Query()
	q.Add("first_name", "Jon")
	q.Add("last_name", "Smith")
	q.Add("age", "34")
	q.Add("children", "Sean,Jessica")
	req.URL.RawQuery = q.Encode()

	if err := schema.Validate(req); err != nil {
		panic(err)
	}
	fmt.Println("HTTP request query validated: success")
}

func RequestQueryFail(schema query.Schema) {
	req := httptest.NewRequest(http.MethodGet, "https://www.query-example.com", nil)
	q := req.URL.Query()
	q.Add("first_name", "Jessica")
	q.Add("last_name", "Doe")
	q.Add("age", "16")
	req.URL.RawQuery = q.Encode()

	err := schema.Validate(req)
	var schemaErr *rerror.SchemaErr
	switch {
	case err == nil:
		panic("error expected for the request query fail")
	case errors.As(err, &schemaErr):
		enc, _ := json.MarshalIndent(schemaErr, "", "  ")
		fmt.Println("HTTP request query validated: failure")
		fmt.Println(string(enc))
	default:
		panic(err)
	}
}
