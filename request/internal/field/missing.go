package field

import "github.com/g8rswimmer/httpx/request/rerror"

func Validate[M1 ~map[string]V1, M2 ~map[string]V2, V1, V2 any](req M1, schema M2) error {
	unknown := []string{}
	for k := range req {
		if _, has := schema[k]; !has {
			unknown = append(unknown, k)
		}
	}
	switch {
	case len(unknown) > 0:
		return &rerror.FieldErr{
			Msg:     "unknown fields are present",
			Unknown: unknown,
		}
	default:
		return nil
	}
}
