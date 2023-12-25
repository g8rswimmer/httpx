package query

import (
	"github.com/g8rswimmer/httpx/request/parameter"
)

type TimeValidator struct {
	parameter.TimeValidator
}

func (p TimeValidator) Validate(value string) error {
	return p.TimeValidator.Validate(value)
}
