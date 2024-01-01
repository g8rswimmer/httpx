package rerror

import "errors"

type SchemaErr struct {
	Msg       string
	Field     *FieldErr     `json:"field_errors,omitempty"`
	Parameter *ParameterErr `json:"parameter_errors,omitempty"`
	Err       string        `json:"error,omitempty"`
}

func (s SchemaErr) Error() string {
	return s.Msg
}

func (s *SchemaErr) Is(target error) bool {
	_, ok := target.(*SchemaErr)
	return ok
}

func SchemaFromError(msg string, err error) error {
	var (
		fieldErr    *FieldErr
		paramterErr *ParameterErr
	)
	switch {
	case err == nil:
		return nil
	case errors.As(err, &fieldErr):
		return &SchemaErr{
			Msg:   msg,
			Field: fieldErr,
		}
	case errors.As(err, &paramterErr):
		if !paramterErr.Has() {
			return nil
		}
		return &SchemaErr{
			Msg:       msg,
			Parameter: paramterErr,
		}
	default:
		return &SchemaErr{
			Msg: msg,
			Err: err.Error(),
		}
	}
}
