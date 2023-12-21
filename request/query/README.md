# Request Query
The query package contians validation around schema and parameters.

## Schema
The query schema is defined by the following structure.  This is used to validate the request query parameters from the key and value.

### Root
| Name             | Type    | Required | JSON Field         | Description                                                                                                                                                                             |
|------------------|---------|----------|--------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Title            | string  | Y        | `title`            | The title of the schema.                                                                                                                                                                |
| Description      | string  | N        | `description`      | Description of the request query parameters                                                                                                                                             |
| Loose Validation | boolean | N        | `loose_validation` | This indicates that schema parameter validation will be "loose". Loose is defined by if there are more parameters than defined in the schema, it will still pass.   Default is `false`. |
| Parameters       | object  | Y        | `parameters`       | This is an object map that defines the parameter and the properties that are assocaited with it.                                                                                        |

### Parameter Properties
| Name                   | Type               | Required                     | JSON Field               | Description                                                                                                                            |
|------------------------|--------------------|------------------------------|--------------------------|----------------------------------------------------------------------------------------------------------------------------------------|
| Description            | string             | N                            | `description`            | Description of the request query parameters                                                                                            |
| Example                | string             | N                            | `example`                | An example of common values for the parameter.                                                                                         |
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
    "loose_validation": false,
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
            "data_type": "boolean"
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
            "inline_array_seperator": ","
        }
    }
}
```
