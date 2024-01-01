package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"

	"github.com/g8rswimmer/httpx/request"
	"github.com/g8rswimmer/httpx/request/rerror"
)

func main() {
	f, err := os.Open("schema.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	schema, err := request.SchemaFromJSON(f)
	if err != nil {
		panic(err)
	}

	RequestSuccess(schema)

	RequestFailureEndpoint(schema)

	RequestFailureQuery(schema)

	RequestFailureBody(schema)
}

func RequestSuccess(schema request.Schema) {
	b := `{
		"first_name": "Gary",
		"last_name": "Doe",
		"alias": ["Gar"],
		"date_of_birth": "1999-05-15",
		"home_address": {
			"street_1": "123 Main St",
			"city": "Springfield",
			"state": "IL"
		},
		"other_addresses": [
			{
				"street_1": "123 Fun Way",
				"city": "Springfield",
				"state": "KY",
				"zip": "41101"
			}
		],
		"phone_number": "+16065551212",
		"married": false,
		"email": "gary.doe@gmail.com"
	}`

	req := httptest.NewRequest(http.MethodPatch, "https://www.query-example.com/schema/example/8cf69907-82f9-4504-8b72-9a608b6381ec", strings.NewReader(b))
	q := req.URL.Query()
	q.Add("bu", "bu_2")
	req.URL.RawQuery = q.Encode()

	if _, err := schema.Validate(req); err != nil {
		panic(err)
	}
	fmt.Println("HTTP request validated: success")
}

func RequestFailureEndpoint(schema request.Schema) {
	b := `{
		"first_name": "Gary",
		"last_name": "Doe",
		"alias": ["Gar"],
		"date_of_birth": "1999-05-15",
		"home_address": {
			"street_1": "123 Main St",
			"city": "Springfield",
			"state": "IL"
		},
		"other_addresses": [
			{
				"street_1": "123 Fun Way",
				"city": "Springfield",
				"state": "KY",
				"zip": "41101"
			}
		],
		"phone_number": "+16065551212",
		"married": false,
		"email": "gary.doe@gmail.com"
	}`

	req := httptest.NewRequest(http.MethodPatch, "https://www.query-example.com/schema/example/8cf69907-82f9-4504-8b72-9a608b6381ecxx", strings.NewReader(b))
	q := req.URL.Query()
	q.Add("bu", "bu_2")
	req.URL.RawQuery = q.Encode()

	_, err := schema.Validate(req)
	var schemaErr *rerror.SchemaErr
	switch {
	case err == nil:
		panic("error expected for the request endpoint fail")
	case errors.As(err, &schemaErr):
		enc, _ := json.MarshalIndent(schemaErr, "", "  ")
		fmt.Println("HTTP request validated: endpoint failure")
		fmt.Println(string(enc))
	default:
		panic(err)
	}
}

func RequestFailureQuery(schema request.Schema) {
	b := `{
		"first_name": "Gary",
		"last_name": "Doe",
		"alias": ["Gar"],
		"date_of_birth": "1999-05-15",
		"home_address": {
			"street_1": "123 Main St",
			"city": "Springfield",
			"state": "IL"
		},
		"other_addresses": [
			{
				"street_1": "123 Fun Way",
				"city": "Springfield",
				"state": "KY",
				"zip": "41101"
			}
		],
		"phone_number": "+16065551212",
		"married": false,
		"email": "gary.doe@gmail.com"
	}`

	req := httptest.NewRequest(http.MethodPatch, "https://www.query-example.com/schema/example/8cf69907-82f9-4504-8b72-9a608b6381ec", strings.NewReader(b))
	q := req.URL.Query()
	q.Add("bu", "bu_4")
	req.URL.RawQuery = q.Encode()

	_, err := schema.Validate(req)
	var schemaErr *rerror.SchemaErr
	switch {
	case err == nil:
		panic("error expected for the request query fail")
	case errors.As(err, &schemaErr):
		enc, _ := json.MarshalIndent(schemaErr, "", "  ")
		fmt.Println("HTTP request validated: query failure")
		fmt.Println(string(enc))
	default:
		panic(err)
	}
}

func RequestFailureBody(schema request.Schema) {
	b := `{
		"first_name": "Gary",
		"last_name": "Doe",
		"alias": ["Gar"],
		"date_of_birth": "1999-05-15",
		"home_address": {
			"street_1": "123 Main St",
			"city": "Springfield",
			"state": "IL"
		},
		"other_addresses": [
			{
				"street_1": "123 Fun Way",
				"city": "Springfield"
			}
		],
		"phone_number": "+16065551212",
		"married": false,
		"email": "gary.doe@gmail.com"
	}`

	req := httptest.NewRequest(http.MethodPatch, "https://www.query-example.com/schema/example/8cf69907-82f9-4504-8b72-9a608b6381ec", strings.NewReader(b))
	q := req.URL.Query()
	q.Add("bu", "bu_2")
	req.URL.RawQuery = q.Encode()

	_, err := schema.Validate(req)
	var schemaErr *rerror.SchemaErr
	switch {
	case err == nil:
		panic("error expected for the request body fail")
	case errors.As(err, &schemaErr):
		enc, _ := json.MarshalIndent(schemaErr, "", "  ")
		fmt.Println("HTTP request validated: body failure")
		fmt.Println(string(enc))
	default:
		panic(err)
	}
}
