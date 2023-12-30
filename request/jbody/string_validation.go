package jbody

import (
	"fmt"

	"github.com/g8rswimmer/httpx/request/parameter"
)

type StringValidator struct {
	parameter.StringValidator
}

func (s StringValidator) Validate(value any) error {
	switch v := value.(type) {
	case string:
		return s.StringValidator.Validate(v)
	default:
		return fmt.Errorf("value is not a string [%T]", value)
	}
}

type StringArrayValidator struct {
	parameter.StringArrayValidator
}

func (s StringArrayValidator) Validate(value any) error {
	var strArr []string
	switch arr := value.(type) {
	case []any:
		for _, v := range arr {
			str, ok := v.(string)
			if !ok {
				return fmt.Errorf("value is not a string [%T]", v)
			}
			strArr = append(strArr, str)
		}
	case []string:
		strArr = arr
	default:
		return fmt.Errorf("value is an array [%T]", arr)
	}
	return s.StringArrayValidator.Validate(strArr)
}
