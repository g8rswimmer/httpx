package endpoint

import (
	"testing"

	"github.com/g8rswimmer/httpx/request/parameter"
)

func TestPathVariable_Validate(t *testing.T) {
	type fields struct {
		Validation VariableValidation
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
			name: "success: string",
			fields: fields{
				Validation: VariableValidation{
					String: &StringValidator{
						StringValidator: parameter.StringValidator{
							RegEx: func() *string {
								s := parameter.RegExUUIDv4
								return &s
							}(),
						},
					},
				},
			},
			args: args{
				value: "1bc71479-6bfd-428b-8c83-f05a2c1df1ec",
			},
			wantErr: false,
		},
		{
			name: "success: number",
			fields: fields{
				Validation: VariableValidation{
					Number: &NumberValidator{
						NumberValidator: parameter.NumberValidator{
							Value: func() *float64 {
								n := 34.0
								return &n
							}(),
						},
					},
				},
			},
			args: args{
				value: "34",
			},
			wantErr: false,
		},
		{
			name: "failure: string",
			fields: fields{
				Validation: VariableValidation{
					String: &StringValidator{
						StringValidator: parameter.StringValidator{
							RegEx: func() *string {
								s := parameter.RegExUUIDv4
								return &s
							}(),
						},
					},
				},
			},
			args: args{
				value: "nope",
			},
			wantErr: true,
		},
		{
			name: "failure: number",
			fields: fields{
				Validation: VariableValidation{
					Number: &NumberValidator{
						NumberValidator: parameter.NumberValidator{
							Value: func() *float64 {
								n := 34.0
								return &n
							}(),
						},
					},
				},
			},
			args: args{
				value: "3334",
			},
			wantErr: true,
		},
		{
			name: "no validation",
			fields: fields{
				Validation: VariableValidation{},
			},
			args: args{
				value: "3334",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pv := PathVariable{
				Validation: tt.fields.Validation,
			}
			if err := pv.Validate(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("PathVariable.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSchemaModelPathVariableValidator(t *testing.T) {
	type args struct {
		pathVariable PathVariable
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				pathVariable: PathVariable{
					Validation: VariableValidation{
						String: &StringValidator{},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "failure",
			args: args{
				pathVariable: PathVariable{
					Validation: VariableValidation{
						String: &StringValidator{},
						Number: &NumberValidator{},
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SchemaModelPathVariableValidator(tt.args.pathVariable); (err != nil) != tt.wantErr {
				t.Errorf("SchemaModelPathVariableValidator() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
