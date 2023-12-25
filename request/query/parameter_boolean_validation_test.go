package query

import (
	"testing"

	"github.com/g8rswimmer/httpx/request/parameter"
)

func TestParameterBooleanValidation_Validate_Value(t *testing.T) {
	type fields struct {
		Value *bool
	}
	type args struct {
		value string
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
				Value: func() *bool {
					b := true
					return &b
				}(),
			},
			args: args{
				value: "true",
			},
			wantErr: false,
		},
		{
			name: "failure: value",
			fields: fields{
				Value: func() *bool {
					b := true
					return &b
				}(),
			},
			args: args{
				value: "false",
			},
			wantErr: true,
		},
		{
			name: "failure: parse",
			fields: fields{
				Value: func() *bool {
					b := true
					return &b
				}(),
			},
			args: args{
				value: "not a boolean",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := QueryBooleanValidator{
				BooleanValidator: parameter.BooleanValidator{
					Value: tt.fields.Value,
				},
			}
			if err := p.Validate(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("ParameterBooleanValidation.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
