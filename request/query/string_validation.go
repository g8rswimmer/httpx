package query

import (
	"github.com/g8rswimmer/httpx/request/parameter"
)

type QueryStringValidator struct {
	parameter.StringValidator
}

func (p QueryStringValidator) Validate(value string) error {
	return p.StringValidator.Validate(value)
}
