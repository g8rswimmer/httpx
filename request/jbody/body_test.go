package jbody

import (
	"encoding/json"
	"testing"

	"github.com/g8rswimmer/httpx/request/internal/field"
	"github.com/g8rswimmer/httpx/request/internal/parameter"
)

func TestBody_Validate(t *testing.T) {
	type fields struct {
		Object      *ObjectValidator
		ObjectArray *ObjectArrayValidator
	}
	type args struct {
		body any
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success: object",
			fields: fields{
				Object: &ObjectValidator{
					RequiredFields: field.Required{},
					Parameters: map[string]ParameterProperties{
						"first_name": {
							Validation: ParameterValidation{
								String: &StringValidator{
									StringValidator: parameter.StringValidator{
										Value: func() *string {
											s := "Gary"
											return &s
										}(),
									},
								},
							},
						},
						"last_name": {
							Validation: ParameterValidation{
								String: &StringValidator{
									StringValidator: parameter.StringValidator{
										Value: func() *string {
											s := "Doe"
											return &s
										}(),
									},
								},
							},
						},
						"age": {
							Validation: ParameterValidation{
								Number: &NumberValidator{},
							},
						},
						"address": {
							Validation: ParameterValidation{
								Object: &ObjectValidator{
									RequiredFields: field.Required{},
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
						},
					},
				},
			},
			args: args{
				body: func() any {
					str := `{
						"value": {
							"first_name": "Gary",
							"last_name": "Doe",
							"age": 34,
							"address": {
								"street_1": "123 Main St",
								"city": "Springfield",
								"state": "IL",
								"zip": "12345"
							}
						}
					}`
					var obj map[string]any
					_ = json.Unmarshal([]byte(str), &obj)
					return obj["value"]
				}(),
			},
			wantErr: false,
		},
		{
			name: "success: object array",
			fields: fields{
				ObjectArray: &ObjectArrayValidator{
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
			},
			args: args{
				body: func() any {
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
		{
			name:    "failure: no validation",
			fields:  fields{},
			args:    args{},
			wantErr: true,
		},
		{
			name: "failure: both validation",
			fields: fields{
				Object:      &ObjectValidator{},
				ObjectArray: &ObjectArrayValidator{},
			},
			args:    args{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Body{
				Object:      tt.fields.Object,
				ObjectArray: tt.fields.ObjectArray,
			}
			if err := b.Validate(tt.args.body); (err != nil) != tt.wantErr {
				t.Errorf("Body.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
