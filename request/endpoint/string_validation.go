package endpoint

import "github.com/g8rswimmer/httpx/request/internal/parameter"

type StringValidator struct {
	parameter.StringValidator
}

func (p StringValidator) Validate(value string) error {
	return p.StringValidator.Validate(value)
}
