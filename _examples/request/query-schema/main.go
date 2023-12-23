package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/g8rswimmer/httpx/request/query"
)

func main() {
	fmt.Println("Query schema example")

	schema := &query.Schema{
		Title:       "Query Schema Example",
		Description: "This is an example of the query schema",
		Parameters: map[string]query.ParameterProperties{
			"first_name": {
				Description: "The first name of the person",
				Example:     "Jon",
				DataType:    query.DataTypeString,
			},
			"last_name": {
				Description: "The last name of the person",
				Example:     "Dow",
				DataType:    query.DataTypeString,
			},
			"age": {
				Description: "The age of the person",
				Example:     "42",
				DataType:    query.DataTypeNumber,
			},
			"married": {
				Description: "If the person is married",
				Example:     "true",
				DataType:    query.DataTypeBoolean,
			},
			"children": {
				Description:          "The list of children of the person, optional",
				Example:              "Jan,Steve",
				DataType:             query.DataTypeString,
				InlineArray:          true,
				InlineArraySeperator: ",",
				Optional:             true,
			},
		},
	}

	fmt.Println("Successful Request Query Schema")
	successReq := httptest.NewRequest(http.MethodGet, "https://www.test.com", nil)
	successQuery := successReq.URL.Query()
	successQuery.Add("first_name", "Jon")
	successQuery.Add("last_name", "Dow")
	successQuery.Add("age", "55")
	successQuery.Add("married", "true")
	successQuery.Add("children", "Mike,Jan")
	successReq.URL.RawQuery = successQuery.Encode()

	if err := schema.Validate(successReq); err != nil {
		panic(err)
	}

	fmt.Println("Failure Request Query Schema schema")
	failReq := httptest.NewRequest(http.MethodGet, "https://www.test.com", nil)
	failQuery := failReq.URL.Query()
	failQuery.Add("first_name", "Jon")
	failQuery.Add("last_name", "Dow")
	failQuery.Add("age", "not a number")
	failQuery.Add("married", "not a boolean")
	failReq.URL.RawQuery = failQuery.Encode()

	err := schema.Validate(failReq)
	if err == nil {
		panic("should be an error")
	}
	var schemaErr *query.SchemaError
	if !errors.As(err, &schemaErr) {
		panic("error not schema err")
	}
	enc, err := json.MarshalIndent(schemaErr, "", "  ")
	if err != nil {
		panic("can't marshal schema error")
	}
	fmt.Println(string(enc))
}
