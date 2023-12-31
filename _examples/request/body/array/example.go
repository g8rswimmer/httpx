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

	if err := schema.Validate(req); err != nil {
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

	err := schema.Validate(req)
	if err == nil {
		panic(err)
	}
	fmt.Println(err.Error())
	fmt.Println("HTTP request body validated: failed")
}
