package parameter

import (
	"errors"
	"fmt"
	"time"
)

type TimeValidator struct {
	Format string  `json:"format"`
	Value  *string `json:"value"`
	Before *string `json:"before"`
	After  *string `json:"after"`
}

func (p TimeValidator) Validate(value string) error {
	if len(p.Format) == 0 {
		return errors.New("time format is required")
	}
	t, err := time.Parse(p.Format, value)
	if err != nil {
		return fmt.Errorf("value [%s] parsing err: %w", value, err)
	}
	if p.Value != nil && *p.Value != value {
		return fmt.Errorf("value [%s] does not match expected [%s]", value, *p.Value)
	}
	if p.Before != nil {
		b, err := time.Parse(p.Format, *p.Before)
		if err != nil {
			return fmt.Errorf("value [%s] parsing err: %w", *p.Before, err)
		}
		if !t.Before(b) {
			return fmt.Errorf("value [%s] is not before [%s]", value, *p.Before)
		}
	}
	if p.After != nil {
		b, err := time.Parse(p.Format, *p.After)
		if err != nil {
			return fmt.Errorf("value [%s] parsing err: %w", *p.After, err)
		}
		if !t.After(b) {
			return fmt.Errorf("value [%s] is not after [%s]", value, *p.After)
		}
	}
	return nil
}
