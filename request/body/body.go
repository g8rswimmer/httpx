package body

import "fmt"

type Body struct {
	Object      *ObjectValidator      `json:"object"`
	ObjectArray *ObjectArrayValidator `json:"object_array"`
}

func (b Body) Validate(body any) error {
	switch {
	case b.Object != nil && b.ObjectArray != nil:
		return fmt.Errorf("body validation can not be an object AND object array")
	case b.Object != nil:
		return b.Object.Validate(body)
	case b.ObjectArray != nil:
		return b.ObjectArray.Validate(body)
	default:
		return fmt.Errorf("body validation not an object or object array [%T]", body)
	}
}
