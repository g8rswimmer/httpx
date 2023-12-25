package field

import "fmt"

type Required struct {
	OneOf   [][]string          `json:"one_of"`
	Present map[string][]string `json:"present"`
}

func (r Required) Validate(fields map[string]struct{}) error {
	if err := findOneOf(r.OneOf, fields); err != nil {
		return err
	}
	if err := findIfPresent(r.Present, fields); err != nil {
		return err
	}
	return nil
}

func find(required []string, fields map[string]struct{}) error {
	for _, r := range required {
		if _, has := fields[r]; !has {
			return fmt.Errorf("[%s] is required", r)
		}
	}
	return nil
}

func findOneOf(oneOf [][]string, fields map[string]struct{}) error {
	if len(oneOf) == 0 {
		return nil
	}
	for _, required := range oneOf {
		if err := find(required, fields); err == nil {
			return nil
		}
	}
	return fmt.Errorf("one of the combinations are required %v", oneOf)
}

func findIfPresent(present map[string][]string, fields map[string]struct{}) error {
	if len(present) == 0 {
		return nil
	}
	for field, required := range present {
		if _, has := fields[field]; !has {
			continue
		}
		if err := find(required, fields); err != nil {
			return fmt.Errorf("if [%s] is present required fields %v", field, required)
		}
	}
	return nil
}
