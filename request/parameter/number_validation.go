package parameter

import (
	"errors"
	"fmt"
)

type NumberValidator struct {
	Value *float64  `json:"value"`
	Min   *float64  `json:"min"`
	Max   *float64  `json:"max"`
	OneOf []float64 `json:"one_of"`
}

func (p NumberValidator) Validate(num float64) error {
	if p.Value != nil && num != *p.Value {
		return fmt.Errorf("value [%f] does not equal %f", num, *p.Value)
	}
	if p.Min != nil && num < *p.Min {
		return fmt.Errorf("value [%f] is less than %f", num, *p.Min)
	}
	if p.Max != nil && num > *p.Max {
		return fmt.Errorf("value [%f] is greater than %f", num, *p.Max)
	}
	if len(p.OneOf) == 0 {
		return nil
	}
	for _, n := range p.OneOf {
		if num == n {
			return nil
		}
	}
	return fmt.Errorf("value [%f] not in %v", num, p.OneOf)
}

type NumberArrayValidator struct {
	Values  []float64 `json:"values"`
	Min     *float64  `json:"min"`
	Max     *float64  `json:"max"`
	Present []float64 `json:"present"`
}

func (n NumberArrayValidator) Validate(nums []float64) error {
	if len(n.Values) > 0 {
		if len(n.Values) != len(nums) {
			return errors.New("validator values lenght must match values length")
		}
		for i := range n.Values {
			if n.Values[i] != nums[i] {
				return fmt.Errorf("value [%f] does not equal %f", nums[i], n.Values[i])
			}
		}
	}
	for _, num := range nums {
		if n.Min != nil && num < *n.Min {
			return fmt.Errorf("value [%f] is less than %f", num, *n.Min)
		}
		if n.Max != nil && num > *n.Max {
			return fmt.Errorf("value [%f] is greater than %f", num, *n.Max)
		}
	}
	if len(n.Present) == 0 {
		return nil
	}
	nset := map[float64]struct{}{}
	for _, n := range nums {
		nset[n] = struct{}{}
	}
	for _, p := range n.Present {
		if _, has := nset[p]; !has {
			return fmt.Errorf("value [%f] not in %v", p, nums)
		}
	}
	return nil
}
