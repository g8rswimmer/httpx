package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"

	"github.com/g8rswimmer/httpx/request/jbody"
)

func main() {
	f, err := os.Open("object_schema.json")
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
	b := `{
		"id": "8b4a60d8-c203-460f-92eb-82646c93d792",
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

	req := httptest.NewRequest(http.MethodPost, "http://www.test.com/schema", strings.NewReader(b))

	if err := schema.Validate(req); err != nil {
		panic(err)
	}

	fmt.Println("HTTP request body validated")
}

func RequestJBodyFail(schema jbody.Schema) {
	b := `{
		"id": "not an id",
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

	req := httptest.NewRequest(http.MethodPost, "http://www.test.com/schema", strings.NewReader(b))

	err := schema.Validate(req)
	if err == nil {
		panic(err)
	}
	fmt.Println(err.Error())
	fmt.Println("HTTP request body validated: failed")
}
