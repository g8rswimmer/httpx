package jbody

import (
	"fmt"
)

type ObjectArrayValidator struct {
	Object ObjectValidator
}

func (o ObjectArrayValidator) Validate(value any) error {
	var objs []map[string]any
	switch v := value.(type) {
	case []any:
		for _, o := range v {
			obj, ok := o.(map[string]any)
			if !ok {
				return fmt.Errorf("value is not an object [%T]", o)
			}
			objs = append(objs, obj)
		}
	case []map[string]any:
		objs = v
	default:
		return fmt.Errorf("value is not an object array [%T]", value)
	}

	for i, obj := range objs {
		if err := o.Object.Validate(obj); err != nil {
			return fmt.Errorf("object array error idx [%d]: %w", i, err)
		}
	}
	return nil
}
