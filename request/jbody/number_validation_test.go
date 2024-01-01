package jbody

import (
	"encoding/json"
	"testing"

	"github.com/g8rswimmer/httpx/request/internal/parameter"
)

func TestNumberValidator_Validate(t *testing.T) {
	type fields struct {
		NumberValidator parameter.NumberValidator
	}
	type args struct {
		value any
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
				NumberValidator: parameter.NumberValidator{
					Value: func() *float64 {
						n := 42.0
						return &n
					}(),
				},
			},
			args: args{
				value: func() any {
					str := `{"value":42}`
					var obj map[string]any
					_ = json.Unmarshal([]byte(str), &obj)
					return obj["value"]
				}(),
			},
			wantErr: false,
		},
		{
			name: "failure",
			fields: fields{
				NumberValidator: parameter.NumberValidator{
					Value: func() *float64 {
						n := 42.0
						return &n
					}(),
				},
			},
			args: args{
				value: func() any {
					str := `{"value":"not a number"}`
					var obj map[string]any
					_ = json.Unmarshal([]byte(str), &obj)
					return obj["value"]
				}(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := NumberValidator{
				NumberValidator: tt.fields.NumberValidator,
			}
			if err := n.Validate(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("NumberValidator.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNumberArrayValidator_Validate(t *testing.T) {
	type fields struct {
		NumberArrayValidator parameter.NumberArrayValidator
	}
	type args struct {
		value any
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
				NumberArrayValidator: parameter.NumberArrayValidator{
					Values: []float64{65, 88},
				},
			},
			args: args{
				value: func() any {
					str := `{"value":[65,88]}`
					var obj map[string]any
					_ = json.Unmarshal([]byte(str), &obj)
					return obj["value"]
				}(),
			},
			wantErr: false,
		},
		{
			name: "failure",
			fields: fields{
				NumberArrayValidator: parameter.NumberArrayValidator{
					Values: []float64{65, 88},
				},
			},
			args: args{
				value: func() any {
					str := `{"value":["sixt-five","eighty-eight"]}`
					var obj map[string]any
					_ = json.Unmarshal([]byte(str), &obj)
					return obj["value"]
				}(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := NumberArrayValidator{
				NumberArrayValidator: tt.fields.NumberArrayValidator,
			}
			if err := n.Validate(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("NumberArrayValidator.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
