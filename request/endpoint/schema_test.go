package endpoint

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/g8rswimmer/httpx/request/internal/parameter"
)

func TestSchema_Validate(t *testing.T) {
	type fields struct {
		Title         string
		Description   string
		Method        string
		Endpoint      string
		PathVariables map[string]PathVariable
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
			name: "success: equals",
			fields: fields{
				Title:       "test",
				Description: "testing endpoint schema valiation",
				Method:      http.MethodGet,
				Endpoint:    "/success/test",
			},
			args: args{
				req: func() *http.Request {
					r := httptest.NewRequest(http.MethodGet, "https://www.schema.com/success/test", nil)
					return r
				}(),
			},
			wantErr: false,
		},
		{
			name: "success: variable",
			fields: fields{
				Title:       "test",
				Description: "testing endpoint schema valiation",
				Method:      http.MethodGet,
				Endpoint:    "/success/test/{id}",
				PathVariables: map[string]PathVariable{
					"{id}": {
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
				},
			},
			args: args{
				req: func() *http.Request {
					r := httptest.NewRequest(http.MethodGet, "https://www.schema.com/success/test/8cf69907-82f9-4504-8b72-9a608b6381ec", nil)
					return r
				}(),
			},
			wantErr: false,
		},
		{
			name: "failure: equals",
			fields: fields{
				Title:       "test",
				Description: "testing endpoint schema valiation",
				Method:      http.MethodGet,
				Endpoint:    "/success/test",
			},
			args: args{
				req: func() *http.Request {
					r := httptest.NewRequest(http.MethodGet, "https://www.schema.com/failure/test", nil)
					return r
				}(),
			},
			wantErr: true,
		},
		{
			name: "failure: variable",
			fields: fields{
				Title:       "test",
				Description: "testing endpoint schema valiation",
				Method:      http.MethodGet,
				Endpoint:    "/success/test/{id}",
				PathVariables: map[string]PathVariable{
					"{id}": {
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
				},
			},
			args: args{
				req: func() *http.Request {
					r := httptest.NewRequest(http.MethodGet, "https://www.schema.com/success/test/no-id", nil)
					return r
				}(),
			},
			wantErr: true,
		},
		{
			name: "failure: method",
			fields: fields{
				Title:       "test",
				Description: "testing endpoint schema valiation",
				Method:      http.MethodGet,
				Endpoint:    "/success/test",
			},
			args: args{
				req: func() *http.Request {
					r := httptest.NewRequest(http.MethodPost, "https://www.schema.com/success/test", nil)
					return r
				}(),
			},
			wantErr: true,
		},
		{
			name: "failure: path length",
			fields: fields{
				Title:       "test",
				Description: "testing endpoint schema valiation",
				Method:      http.MethodGet,
				Endpoint:    "/success/test",
			},
			args: args{
				req: func() *http.Request {
					r := httptest.NewRequest(http.MethodGet, "https://www.schema.com/success/test/extra", nil)
					return r
				}(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Schema{
				Title:         tt.fields.Title,
				Description:   tt.fields.Description,
				Method:        tt.fields.Method,
				Endpoint:      tt.fields.Endpoint,
				PathVariables: tt.fields.PathVariables,
			}
			if err := s.Validate(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("Schema.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSchemaModelValidator(t *testing.T) {
	type args struct {
		schema Schema
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				schema: Schema{
					Method:   http.MethodGet,
					Endpoint: "/test/{id}",
					PathVariables: map[string]PathVariable{
						"{id}": {
							Validation: VariableValidation{
								String: &StringValidator{},
							},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "failure: no method",
			args: args{
				schema: Schema{
					Endpoint: "/test/{id}",
					PathVariables: map[string]PathVariable{
						"{id}": {
							Validation: VariableValidation{
								String: &StringValidator{},
							},
						},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "failure: path variables",
			args: args{
				schema: Schema{
					Method:   http.MethodGet,
					Endpoint: "/test/{id}",
					PathVariables: map[string]PathVariable{
						"{id}": {
							Validation: VariableValidation{},
						},
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SchemaModelValidator(tt.args.schema); (err != nil) != tt.wantErr {
				t.Errorf("SchemaModelValidator() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
