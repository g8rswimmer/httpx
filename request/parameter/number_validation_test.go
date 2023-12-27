package parameter

import (
	"testing"
)

func TestParameterDataNumberValidation_Validate_Value(t *testing.T) {

	type fields struct {
		Value *float64
	}
	type args struct {
		value float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				Value: func() *float64 {
					n := 34.0
					return &n
				}(),
			},
			args: args{
				value: 34.0,
			},
			wantErr: false,
		},
		{
			name: "failure",
			fields: fields{
				Value: func() *float64 {
					n := 64.0
					return &n
				}(),
			},
			args: args{
				value: 34,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NumberValidator{
				Value: tt.fields.Value,
			}
			if err := p.Validate(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("ParameterDataNumberValidation.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParameterDataNumberValidation_Validate_Min_Max(t *testing.T) {
	type fields struct {
		Min *float64
		Max *float64
	}
	type args struct {
		value float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success: min",
			fields: fields{
				Min: func() *float64 {
					n := 10.0
					return &n
				}(),
			},
			args: args{
				value: 12,
			},
			wantErr: false,
		},
		{
			name: "success: max",
			fields: fields{
				Max: func() *float64 {
					n := 10.0
					return &n
				}(),
			},
			args: args{
				value: 6,
			},
			wantErr: false,
		},
		{
			name: "success: min and max",
			fields: fields{
				Min: func() *float64 {
					n := 1.0
					return &n
				}(),
				Max: func() *float64 {
					n := 10.0
					return &n
				}(),
			},
			args: args{
				value: 6,
			},
			wantErr: false,
		},
		{
			name: "failure: min",
			fields: fields{
				Min: func() *float64 {
					n := 10.0
					return &n
				}(),
			},
			args: args{
				value: 9,
			},
			wantErr: true,
		},
		{
			name: "failure: max",
			fields: fields{
				Max: func() *float64 {
					n := 10.0
					return &n
				}(),
			},
			args: args{
				value: 16,
			},
			wantErr: true,
		},
		{
			name: "failure: min and max",
			fields: fields{
				Min: func() *float64 {
					n := 1.0
					return &n
				}(),
				Max: func() *float64 {
					n := 10.0
					return &n
				}(),
			},
			args: args{
				value: 116,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NumberValidator{
				Min: tt.fields.Min,
				Max: tt.fields.Max,
			}
			if err := p.Validate(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("ParameterDataNumberValidation.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParameterDataNumberValidation_Validate_OneOf(t *testing.T) {
	type fields struct {
		OneOf []float64
	}
	type args struct {
		value float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				OneOf: []float64{10.0, 34.0},
			},
			args: args{
				value: 34,
			},
			wantErr: false,
		},
		{
			name: "failure",
			fields: fields{
				OneOf: []float64{10.0, 34.0},
			},
			args: args{
				value: 38,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NumberValidator{
				OneOf: tt.fields.OneOf,
			}
			if err := p.Validate(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("ParameterDataNumberValidation.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNumberArrayValidator_Validate_Values(t *testing.T) {
	type fields struct {
		Values []float64
	}
	type args struct {
		nums []float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				Values: []float64{12, 34, 98},
			},
			args: args{
				nums: []float64{12, 34, 98},
			},
			wantErr: false,
		},
		{
			name: "failure: value",
			fields: fields{
				Values: []float64{12, 34, 98},
			},
			args: args{
				nums: []float64{12, 111, 98},
			},
			wantErr: true,
		},
		{
			name: "failure: length",
			fields: fields{
				Values: []float64{12, 34, 98, 1},
			},
			args: args{
				nums: []float64{12, 34, 98},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := NumberArrayValidator{
				Values: tt.fields.Values,
			}
			if err := n.Validate(tt.args.nums); (err != nil) != tt.wantErr {
				t.Errorf("NumberArrayValidator.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNumberArrayValidator_Validate_Min_Max(t *testing.T) {
	type fields struct {
		Min *float64
		Max *float64
	}
	type args struct {
		nums []float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				Min: func() *float64 {
					m := 10.0
					return &m
				}(),
				Max: func() *float64 {
					m := 52.0
					return &m
				}(),
			},
			args: args{
				nums: []float64{11, 34, 51},
			},
			wantErr: false,
		},
		{
			name: "failure: min",
			fields: fields{
				Min: func() *float64 {
					m := 10.0
					return &m
				}(),
				Max: func() *float64 {
					m := 52.0
					return &m
				}(),
			},
			args: args{
				nums: []float64{11, 34, 7, 51},
			},
			wantErr: true,
		},
		{
			name: "failure: max",
			fields: fields{
				Min: func() *float64 {
					m := 10.0
					return &m
				}(),
				Max: func() *float64 {
					m := 52.0
					return &m
				}(),
			},
			args: args{
				nums: []float64{11, 34, 117, 51},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := NumberArrayValidator{
				Min: tt.fields.Min,
				Max: tt.fields.Max,
			}
			if err := n.Validate(tt.args.nums); (err != nil) != tt.wantErr {
				t.Errorf("NumberArrayValidator.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNumberArrayValidator_Validate_Present(t *testing.T) {
	type fields struct {
		Present []float64
	}
	type args struct {
		nums []float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				Present: []float64{10, 34},
			},
			args: args{
				nums: []float64{9, 22, 34, 76, 10},
			},
			wantErr: false,
		},
		{
			name: "failure",
			fields: fields{
				Present: []float64{10, 34},
			},
			args: args{
				nums: []float64{9, 22, 5534, 76, 10},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := NumberArrayValidator{
				Present: tt.fields.Present,
			}
			if err := n.Validate(tt.args.nums); (err != nil) != tt.wantErr {
				t.Errorf("NumberArrayValidator.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
