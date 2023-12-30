package body

import (
	"encoding/json"
	"testing"

	"github.com/g8rswimmer/httpx/request/field"
	"github.com/g8rswimmer/httpx/request/parameter"
)

func TestObjectArrayValidator_Validate_Success(t *testing.T) {
	type fields struct {
		Object ObjectValidator
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
				Object: ObjectValidator{
					RequiredFields: field.Required{
						OneOf: [][]string{{"street_1", "zip"}, {"street_1", "city", "state"}},
					},
					Parameters: map[string]ParameterProperties{
						"street_1": {
							Validation: ParameterValidation{
								String: &StringValidator{},
							},
						},
						"street_2": {
							Validation: ParameterValidation{
								String: &StringValidator{},
							},
						},
						"city": {
							Validation: ParameterValidation{
								String: &StringValidator{},
							},
						},
						"state": {
							Validation: ParameterValidation{
								String: &StringValidator{
									StringValidator: parameter.StringValidator{
										RegEx: func() *string {
											r := "^[A-Za-z]{2}$"
											return &r
										}(),
									},
								},
							},
						},
						"zip": {
							Validation: ParameterValidation{
								String: &StringValidator{
									StringValidator: parameter.StringValidator{
										RegEx: func() *string {
											r := "^[0-9]{5}$"
											return &r
										}(),
									},
								},
							},
						},
					},
				},
			},
			args: args{
				value: func() any {
					str := `{
						"value": [
							{
								"street_1": "123 Main St",
								"zip": "12345"
							},
							{
								"street_1": "123 Main St",
								"city": "Springfield",
								"state": "IL"
							}
						]
					}`
					var obj map[string]any
					_ = json.Unmarshal([]byte(str), &obj)
					return obj["value"]
				}(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := ObjectArrayValidator{
				Object: tt.fields.Object,
			}
			if err := o.Validate(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("ObjectArrayValidator.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestObjectArrayValidator_Validate_Failure(t *testing.T) {
	type fields struct {
		Object ObjectValidator
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
			name:   "failure: not object array",
			fields: fields{},
			args: args{
				value: func() any {
					str := `{"value":"some value"}`
					var obj map[string]any
					_ = json.Unmarshal([]byte(str), &obj)
					return obj["value"]
				}(),
			},
			wantErr: true,
		},
		{
			name: "failure: object validation",
			fields: fields{
				Object: ObjectValidator{
					RequiredFields: field.Required{
						OneOf: [][]string{{"street_1", "zip"}, {"street_1", "city", "state"}},
					},
					Parameters: map[string]ParameterProperties{
						"street_1": {
							Validation: ParameterValidation{
								String: &StringValidator{},
							},
						},
						"street_2": {
							Validation: ParameterValidation{
								String: &StringValidator{},
							},
						},
						"city": {
							Validation: ParameterValidation{
								String: &StringValidator{},
							},
						},
						"state": {
							Validation: ParameterValidation{
								String: &StringValidator{
									StringValidator: parameter.StringValidator{
										RegEx: func() *string {
											r := "^[A-Za-z]{2}$"
											return &r
										}(),
									},
								},
							},
						},
						"zip": {
							Validation: ParameterValidation{
								String: &StringValidator{
									StringValidator: parameter.StringValidator{
										RegEx: func() *string {
											r := "^[0-9]{5}$"
											return &r
										}(),
									},
								},
							},
						},
					},
				},
			},
			args: args{
				value: func() any {
					str := `{
						"value": [
							{
								"street_1": "123 Main St",
								"zip": "12345"
							},
							{
								"street_1": "123 Main St",
								"state": "IL"
							}
						]
					}`
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
			o := ObjectArrayValidator{
				Object: tt.fields.Object,
			}
			if err := o.Validate(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("ObjectArrayValidator.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
