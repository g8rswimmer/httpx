package jbody

import (
	"fmt"

	"github.com/g8rswimmer/httpx/request/internal/parameter"
)

type TimeValidator struct {
	parameter.TimeValidator
}

func (t TimeValidator) Validate(value any) error {
	switch v := value.(type) {
	case string:
		return t.TimeValidator.Validate(v)
	default:
		return fmt.Errorf("value is not a string [%T]", value)
	}
}

type TimeArrayValidator struct {
	parameter.TimeArrayValidator
}

func (s TimeArrayValidator) Validate(value any) error {
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
	return s.TimeArrayValidator.Validate(strArr)
}
