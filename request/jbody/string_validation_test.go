package jbody

import (
	"encoding/json"
	"testing"

	"github.com/g8rswimmer/httpx/request/internal/parameter"
)

func TestStringValidator_Validate(t *testing.T) {
	type fields struct {
		StringValidator parameter.StringValidator
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
				StringValidator: parameter.StringValidator{
					Value: func() *string {
						s := "some value"
						return &s
					}(),
				},
			},
			args: args{
				value: func() any {
					str := `{"value":"some value"}`
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
				StringValidator: parameter.StringValidator{
					Value: func() *string {
						s := "some value"
						return &s
					}(),
				},
			},
			args: args{
				value: func() any {
					str := `{"value":43}`
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
			s := StringValidator{
				StringValidator: tt.fields.StringValidator,
			}
			if err := s.Validate(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("StringValidator.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStringArrayValidator_Validate(t *testing.T) {
	type fields struct {
		StringArrayValidator parameter.StringArrayValidator
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
				StringArrayValidator: parameter.StringArrayValidator{
					Values: []string{"value1", "value2"},
				},
			},
			args: args{
				value: func() any {
					str := `{"value":["value1","value2"]}`
					var obj map[string]any
					_ = json.Unmarshal([]byte(str), &obj)
					return obj["value"]
				}(),
			},
			wantErr: false,
		},
		{
			name: "success",
			fields: fields{
				StringArrayValidator: parameter.StringArrayValidator{
					Values: []string{"value1", "value2"},
				},
			},
			args: args{
				value: func() any {
					str := `{"value":[10,88]}`
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
			s := StringArrayValidator{
				StringArrayValidator: tt.fields.StringArrayValidator,
			}
			if err := s.Validate(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("StringArrayValidator.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
