package jbody

import (
	"encoding/json"
	"testing"

	"github.com/g8rswimmer/httpx/request/field"
	"github.com/g8rswimmer/httpx/request/parameter"
)

func TestObjectValidator_Validate_Success(t *testing.T) {
	type fields struct {
		RequiredFields field.Required
		Parameters     map[string]ParameterProperties
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
			args: args{
				value: func() any {
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
			name: "success: required",
			fields: fields{
				RequiredFields: field.Required{
					OneOf: [][]string{{"first_name", "last_name"}},
				},
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
				},
			},
			args: args{
				value: func() any {
					str := `{
						"value": {
							"first_name": "Gary",
							"last_name": "Doe",
							"address": {
								"street_1": "123 Main St",
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := ObjectValidator{
				RequiredFields: tt.fields.RequiredFields,
				Parameters:     tt.fields.Parameters,
			}
			if err := o.Validate(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("ObjectValidator.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestObjectValidator_Validate_Failure(t *testing.T) {
	type fields struct {
		RequiredFields field.Required
		Parameters     map[string]ParameterProperties
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
			name:   "not an object",
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
			name: "missing required fields",
			fields: fields{
				RequiredFields: field.Required{
					OneOf: [][]string{{"last_name", "first_name"}},
				},
			},
			args: args{
				value: func() any {
					str := `{
						"value": {
							"first_name": "Gary",
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
			wantErr: true,
		},
		{
			name: "extra fields",
			fields: fields{
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
				},
			},
			args: args{
				value: func() any {
					str := `{
						"value": {
							"first_name": "Gary",
							"last_name": "Doe",
							"age": 34
						}
					}`
					var obj map[string]any
					_ = json.Unmarshal([]byte(str), &obj)
					return obj["value"]
				}(),
			},
			wantErr: true,
		},
		{
			name: "no validation",
			fields: fields{
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
						Validation: ParameterValidation{},
					},
				},
			},
			args: args{
				value: func() any {
					str := `{
						"value": {
							"first_name": "Gary",
							"last_name": "Doe"
						}
					}`
					var obj map[string]any
					_ = json.Unmarshal([]byte(str), &obj)
					return obj["value"]
				}(),
			},
			wantErr: true,
		},
		{
			name: "too many validators",
			fields: fields{
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
							Number: &NumberValidator{},
						},
					},
				},
			},
			args: args{
				value: func() any {
					str := `{
						"value": {
							"first_name": "Gary",
							"last_name": "Doe"
						}
					}`
					var obj map[string]any
					_ = json.Unmarshal([]byte(str), &obj)
					return obj["value"]
				}(),
			},
			wantErr: true,
		},
		{
			name: "no match",
			fields: fields{
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
				},
			},
			args: args{
				value: func() any {
					str := `{
						"value": {
							"first_name": "Gary",
							"last_name": "Nope"
						}
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
			o := ObjectValidator{
				RequiredFields: tt.fields.RequiredFields,
				Parameters:     tt.fields.Parameters,
			}
			if err := o.Validate(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("ObjectValidator.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
