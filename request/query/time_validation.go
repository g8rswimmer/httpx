package query

import (
	"github.com/g8rswimmer/httpx/request/parameter"
)

type QueryTimeValidator struct {
	parameter.TimeValidator
}

func (p QueryTimeValidator) Validate(value string) error {
	return p.TimeValidator.Validate(value)
}
