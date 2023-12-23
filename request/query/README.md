# Request Query
The query package contians validation around schema and parameters.

## Schema
The query schema is defined by the following structure.  This is used to validate the request query parameters from the key and value.

### Root
| Name             | Type    | Required | JSON Field         | Description                                                                                                                                                                             |
|------------------|---------|----------|--------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Title            | string  | Y        | `title`            | The title of the schema.                                                                                                                                                                |
| Description      | string  | N        | `description`      | Description of the request query parameters                                                                                                                                             |
| Parameters       | object  | Y        | `parameters`       | This is an object map that defines the parameter and the properties that are assocaited with it.                                                                                        |

### Parameter Properties
| Name                   | Type               | Required                     | JSON Field               | Description                                                                                                                            |
|------------------------|--------------------|------------------------------|--------------------------|----------------------------------------------------------------------------------------------------------------------------------------|
| Description            | string             | N                            | `description`            | Description of the request query parameters                                                                                            |
| Example                | string             | N                            | `example`                | An example of common values for the parameter.                                                                                         |
| Optional                | boolean             | N                            | `optional`                | If the parameter is optional.  If in the query it will be cheked, but will not fail if not |
| Data Type              | string enumeration | Y                            | `data_type`              | The data type of the parameter.  While all parameters are "strings" in the query, this is the data type that it would be converted to. |
| Inline Array           | boolean            | N                            | `inline_array`           | Some parameters may use a seperator to represent an array.  This is to indicate that the parameter value.                              |
| Inline Array Seperator | string             | Y, if inline array is `true` | `inline_array_seperator` | The seperator used to seperate the array ellements.                                                                                    |

#### Data Types
The following are supported data types for the parameter properties:
  *  `string`
  *  `number`
  *  `boolean`

### JSON Example
```json
{
    "title": "Example Query Schema",
    "description": "This is an example of the query schema for the following query: ?last_name=Doe&first_name=John&married=true&age=34,children=David,Susan",
    "properties": {
        "first_name": {
            "description": "First Name of the person",
            "example": "John",
            "data_type": "string"
        },
        "last_name": {
            "description": "Last Name of the person",
            "example": "Doe",
            "data_type": "string"
        },
        "married": {
            "description": "Is the person married",
            "data_type": "boolean",
            "optional": true
        },
        "age": {
            "description": "The age of the person",
            "example": "39",
            "data_type": "number"
        },
        "children": {
            "description": "The children of the person",
            "example": "Drew,Erin",
            "data_type": "string",
            "inline_array": true,
            "inline_array_seperator": ",",
            "optional": true
        }
    }
}
```
## Code Example
```go
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
```
This example can be found [here][../../_examples/request/query-schema/main.go]