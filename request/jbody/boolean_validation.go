package jbody

import (
	"fmt"

	"github.com/g8rswimmer/httpx/request/internal/parameter"
)

type BooleanValidator struct {
	parameter.BooleanValidator
}

func (b BooleanValidator) Validate(value any) error {
	switch v := value.(type) {
	case bool:
		return b.BooleanValidator.Validate(v)
	default:
		return fmt.Errorf("value is not a boolean [%T]", value)
	}
}
