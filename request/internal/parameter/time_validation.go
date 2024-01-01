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

type TimeArrayValidator struct {
	Format string   `json:"format"`
	Values []string `json:"values"`
	Before *string  `json:"before"`
	After  *string  `json:"after"`
}

func (tav TimeArrayValidator) Validate(values []string) error {
	if len(tav.Format) == 0 {
		return errors.New("time format is required")
	}
	ts := make([]time.Time, len(values))
	for i, value := range values {
		t, err := time.Parse(tav.Format, value)
		if err != nil {
			return fmt.Errorf("value [%s] parsing err: %w", value, err)
		}
		ts[i] = t
	}
	if len(tav.Values) > 0 {
		if len(tav.Values) != len(values) {
			return errors.New("validator values lenght must match values length")
		}
		for i := range tav.Values {
			if tav.Values[i] != values[i] {
				return fmt.Errorf("value [%s] does not equal %s", values[i], tav.Values[i])
			}
		}
	}
	for _, t := range ts {
		if tav.Before != nil {
			b, err := time.Parse(tav.Format, *tav.Before)
			if err != nil {
				return fmt.Errorf("value [%s] parsing err: %w", *tav.Before, err)
			}
			if !t.Before(b) {
				return fmt.Errorf("value [%s] is not before [%s]", t, *tav.Before)
			}
		}
		if tav.After != nil {
			b, err := time.Parse(tav.Format, *tav.After)
			if err != nil {
				return fmt.Errorf("value [%s] parsing err: %w", *tav.After, err)
			}
			if !t.After(b) {
				return fmt.Errorf("value [%s] is not after [%s]", t, *tav.After)
			}
		}
	}
	return nil
}
