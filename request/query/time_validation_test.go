package query

import (
	"testing"
	"time"

	"github.com/g8rswimmer/httpx/request/parameter"
)

func TestParameterTimeValidation_Validate_Format(t *testing.T) {
	type fields struct {
		Format string
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
				Format: time.RFC3339,
			},
			args: args{
				value: "2023-10-12T07:20:50.52Z",
			},
			wantErr: false,
		},
		{
			name: "failure",
			fields: fields{
				Format: time.RFC822,
			},
			args: args{
				value: "2023-10-12T07:20:50.52Z",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := QueryTimeValidator{
				TimeValidator: parameter.TimeValidator{
					Format: tt.fields.Format,
				},
			}
			if err := p.Validate(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("ParameterTimeValidation.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParameterTimeValidation_Validate_Value(t *testing.T) {
	type fields struct {
		Format string
		Value  *string
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
			name: " success",
			fields: fields{
				Format: time.RFC3339,
				Value: func() *string {
					t := "2023-10-12T07:20:50.52Z"
					return &t
				}(),
			},
			args: args{
				value: "2023-10-12T07:20:50.52Z",
			},
			wantErr: false,
		},
		{
			name: " success",
			fields: fields{
				Format: time.RFC3339,
				Value: func() *string {
					t := "2043-10-12T07:20:50.52Z"
					return &t
				}(),
			},
			args: args{
				value: "2023-10-12T07:20:50.52Z",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := QueryTimeValidator{
				TimeValidator: parameter.TimeValidator{
					Format: tt.fields.Format,
					Value:  tt.fields.Value,
				},
			}
			if err := p.Validate(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("ParameterTimeValidation.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParameterTimeValidation_Validate_Before(t *testing.T) {
	type fields struct {
		Format string
		Before *string
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
			name: " success",
			fields: fields{
				Format: time.RFC3339,
				Before: func() *string {
					t := "2024-10-12T07:20:50.52Z"
					return &t
				}(),
			},
			args: args{
				value: "2023-10-12T07:20:50.52Z",
			},
			wantErr: false,
		},
		{
			name: " failure",
			fields: fields{
				Format: time.RFC3339,
				Before: func() *string {
					t := "2022-10-12T07:20:50.52Z"
					return &t
				}(),
			},
			args: args{
				value: "2023-10-12T07:20:50.52Z",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := QueryTimeValidator{
				TimeValidator: parameter.TimeValidator{
					Format: tt.fields.Format,
					Before: tt.fields.Before,
				},
			}
			if err := p.Validate(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("ParameterTimeValidation.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParameterTimeValidation_Validate_After(t *testing.T) {
	type fields struct {
		Format string
		After  *string
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
			name: " success",
			fields: fields{
				Format: time.RFC3339,
				After: func() *string {
					t := "2022-10-12T07:20:50.52Z"
					return &t
				}(),
			},
			args: args{
				value: "2023-10-12T07:20:50.52Z",
			},
			wantErr: false,
		},
		{
			name: " failure",
			fields: fields{
				Format: time.RFC3339,
				After: func() *string {
					t := "2024-10-12T07:20:50.52Z"
					return &t
				}(),
			},
			args: args{
				value: "2023-10-12T07:20:50.52Z",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := QueryTimeValidator{
				TimeValidator: parameter.TimeValidator{
					Format: tt.fields.Format,
					After:  tt.fields.After,
				},
			}
			if err := p.Validate(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("ParameterTimeValidation.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
