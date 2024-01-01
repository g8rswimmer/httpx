package rerror

type FieldErr struct {
	Msg     string              `json:"message"`
	OneOf   [][]string          `json:"one_of,omitempty"`
	Present map[string][]string `json:"present,omitempty"`
	Unknown []string            `json:"unkown_fields,omitempty"`
}

func (r FieldErr) Error() string {
	return r.Msg
}

func (s *FieldErr) Is(target error) bool {
	_, ok := target.(*FieldErr)
	return ok
}
