package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"

	"github.com/g8rswimmer/httpx/request/jbody"
	"github.com/g8rswimmer/httpx/request/rerror"
)

func main() {
	f, err := os.Open("array_schema.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	schema, err := jbody.SchemaFromJSON(f)
	if err != nil {
		panic(err)
	}

	RequestJBodySuccess(schema)

	RequestJBodyFail(schema)
}

func RequestJBodySuccess(schema jbody.Schema) {
	b := `[
			{
				"street_1": "123 Fun Way",
				"city": "Springfield",
				"state": "KY",
				"zip": "41101"
			}
		]
	}`

	req := httptest.NewRequest(http.MethodPost, "http://www.test.com/schema", strings.NewReader(b))

	if _, err := schema.Validate(req); err != nil {
		panic(err)
	}

	fmt.Println("HTTP request body validated")
}

func RequestJBodyFail(schema jbody.Schema) {
	b := `[
			{
				"street_1": "123 Fun Way",
				"city": "Springfield",
				"state": "KY",
				"zip": "41101"
			},
			{
				"street_1": "123 Fun Way",
				"city": "Springfield"
			}
		]
	}`

	req := httptest.NewRequest(http.MethodPost, "http://www.test.com/schema", strings.NewReader(b))

	_, err := schema.Validate(req)
	var schemaErr *rerror.SchemaErr
	switch {
	case err == nil:
		panic("error expected for the request query fail")
	case errors.As(err, &schemaErr):
		enc, _ := json.MarshalIndent(schemaErr, "", "  ")
		fmt.Println("HTTP request body validated: failure")
		fmt.Println(string(enc))
	default:
		panic(err)
	}
}
