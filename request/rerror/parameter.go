package rerror

type ParameterErr struct {
	Parameters map[string]string `json:"parameters"`
}

func (p *ParameterErr) Add(k string, msg string) {
	p.Parameters[k] = msg
}

func (p ParameterErr) Has() bool {
	return len(p.Parameters) > 0
}

func (p ParameterErr) Error() string {
	return "parameter error"
}

func (s *ParameterErr) Is(target error) bool {
	_, ok := target.(*ParameterErr)
	return ok
}
