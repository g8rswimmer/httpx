package query

import (
	"testing"

	"github.com/g8rswimmer/httpx/request/parameter"
)

func TestParameterStringValidation_Validate_Value(t *testing.T) {
	type fields struct {
		Value *string
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
				Value: func() *string {
					s := "test"
					return &s
				}(),
			},
			args: args{
				value: "test",
			},
			wantErr: false,
		},
		{
			name: "failure",
			fields: fields{
				Value: func() *string {
					s := "test"
					return &s
				}(),
			},
			args: args{
				value: "fail",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := QueryStringValidator{
				StringValidator: parameter.StringValidator{
					Value: tt.fields.Value,
				},
			}
			if err := p.Validate(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("ParameterStringValidation.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParameterStringValidation_Validate_OneOf(t *testing.T) {
	type fields struct {
		OneOf []string
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
				OneOf: []string{"one", "two"},
			},
			args: args{
				value: "two",
			},
			wantErr: false,
		},
		{
			name: "failure",
			fields: fields{
				OneOf: []string{"one", "two"},
			},
			args: args{
				value: "three",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := QueryStringValidator{
				StringValidator: parameter.StringValidator{
					OneOf: tt.fields.OneOf,
				},
			}
			if err := p.Validate(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("ParameterStringValidation.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParameterStringValidation_Validate_RegEx(t *testing.T) {
	type fields struct {
		RegEx *string
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
			name: "success: uuidv4",
			fields: fields{
				RegEx: func() *string {
					s := parameter.RegExUUIDv4
					return &s
				}(),
			},
			args: args{
				value: "ead5b1f7-e8e3-4c65-ba1d-29c98ea5a5b8",
			},
			wantErr: false,
		},
		{
			name: "failure: uuidv4",
			fields: fields{
				RegEx: func() *string {
					s := parameter.RegExUUIDv4
					return &s
				}(),
			},
			args: args{
				value: "eadb1f7-e8e3-4c65-ba1d-29c98ea5a5b8",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := QueryStringValidator{
				StringValidator: parameter.StringValidator{
					RegEx: tt.fields.RegEx,
				},
			}
			if err := p.Validate(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("ParameterStringValidation.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
