package body

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/g8rswimmer/httpx/request/field"
	"github.com/g8rswimmer/httpx/request/parameter"
)

func TestSchema_Validate(t *testing.T) {
	type fields struct {
		Title       string
		Description string
		Body        Body
	}
	type args struct {
		req *http.Request
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
				Body: Body{
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
			},
			args: args{
				req: func() *http.Request {
					b := `{
						"first_name": "Gary",
						"last_name": "Doe",
						"age": 34,
						"address": {
							"street_1": "123 Main St",
							"city": "Springfield",
							"state": "IL",
							"zip": "12345"
						}
					}`
					return httptest.NewRequest(http.MethodPost, "https:\\www.test.this", strings.NewReader(b))
				}(),
			},
			wantErr: false,
		},
		{
			name: "success: object array",
			fields: fields{
				Body: Body{
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
			},
			args: args{
				req: func() *http.Request {
					b := `[
						{
							"street_1": "123 Main St",
							"zip": "12345"
						},
						{
							"street_1": "123 Main St",
							"city": "Springfield",
							"state": "IL"
						}
					]`
					return httptest.NewRequest(http.MethodPost, "https:\\www.test.this", strings.NewReader(b))
				}(),
			},
			wantErr: false,
		},
		{
			name: "failure: validation",
			fields: fields{
				Body: Body{
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
			},
			args: args{
				req: func() *http.Request {
					b := `{
						"first_name": "Gary",
						"last_name": "Doe",
						"age": "34",
						"address": {
							"street_1": "123 Main St",
							"city": "Springfield",
							"state": "IL",
							"zip": "12345"
						}
					}`
					return httptest.NewRequest(http.MethodPost, "https:\\www.test.this", strings.NewReader(b))
				}(),
			},
			wantErr: true,
		},
		{
			name:   "failure: json",
			fields: fields{},
			args: args{
				req: func() *http.Request {
					b := `{
						"first_name": "Gary",
						"last_name": "Doe",
						"age": "34,
						"address": {
							"street_1": "123 Main St",
							"city": "Springfield",
							"state": "IL",
							"zip": "12345"
						}
					}`
					return httptest.NewRequest(http.MethodPost, "https:\\www.test.this", strings.NewReader(b))
				}(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Schema{
				Title:       tt.fields.Title,
				Description: tt.fields.Description,
				Body:        tt.fields.Body,
			}
			if err := s.Validate(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("Schema.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
