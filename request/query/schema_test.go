package query

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

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
					Title:       "Test",
					Description: "test example",
					Parameters: map[string]ParameterProperties{
						"param1": {
							Description: "None",
							Example:     "Test",
							Validation: ParameterValidation{
								String: &ParameterStringValidator{},
							},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "fail: no parameters",
			args: args{
				schema: Schema{
					Title:       "Test",
					Description: "test example",
				},
			},
			wantErr: true,
		},
		{
			name: "fail: parameter properties",
			args: args{
				schema: Schema{
					Title:       "Test",
					Description: "test example",
					Parameters: map[string]ParameterProperties{
						"param1": {
							Description: "None",
							Example:     "Test",
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

func TestSchema_Validate(t *testing.T) {
	type fields struct {
		Title       string
		Description string
		Parameters  map[string]ParameterProperties
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
			name: "success",
			fields: fields{
				Title:       "Schema Test",
				Description: "testing validation",
				Parameters: map[string]ParameterProperties{
					"param1": {
						Description: "String",
						Example:     "Test",
						Validation: ParameterValidation{
							String: &ParameterStringValidator{},
						},
					},
					"param2": {
						Description: "Number",
						Example:     "Test",
						Validation: ParameterValidation{
							Number: &ParameterDataNumberValidator{},
						},
					},
					"param3": {
						Description: "Boolean",
						Example:     "Test",
						Validation: ParameterValidation{
							Boolean: &ParameterBooleanValidator{},
						},
					},
					"param4": {
						Description:          "Number Array",
						Example:              "Test",
						InlineArray:          true,
						InlineArraySeperator: ",",
						Validation: ParameterValidation{
							Number: &ParameterDataNumberValidator{},
						},
					},
				},
			},
			args: args{
				req: func() *http.Request {
					r := httptest.NewRequest(http.MethodGet, "https://www.google.com", nil)
					q := r.URL.Query()
					q.Add("param1", "some string")
					q.Add("param2", "42")
					q.Add("param3", "true")
					q.Add("param4", "13,55")
					r.URL.RawQuery = q.Encode()
					return r
				}(),
			},
			wantErr: false,
		},
		{
			name: "success: optional",
			fields: fields{
				Title:       "Schema Test",
				Description: "testing validation",
				Parameters: map[string]ParameterProperties{
					"param1": {
						Description: "String",
						Example:     "Test",
						Validation: ParameterValidation{
							String: &ParameterStringValidator{},
						},
					},
					"param2": {
						Description: "Number",
						Example:     "Test",
						Validation: ParameterValidation{
							Number: &ParameterDataNumberValidator{},
						},
					},
					"param3": {
						Description: "Boolean",
						Example:     "Test",
						Optional:    true,
						Validation: ParameterValidation{
							Boolean: &ParameterBooleanValidator{},
						},
					},
					"param4": {
						Description:          "Number Array",
						Example:              "Test",
						InlineArray:          true,
						InlineArraySeperator: ",",
						Validation: ParameterValidation{
							Number: &ParameterDataNumberValidator{},
						},
					},
				},
			},
			args: args{
				req: func() *http.Request {
					r := httptest.NewRequest(http.MethodGet, "https://www.google.com", nil)
					q := r.URL.Query()
					q.Add("param1", "some string")
					q.Add("param2", "42")
					q.Add("param4", "13,55")
					r.URL.RawQuery = q.Encode()
					return r
				}(),
			},
			wantErr: false,
		},
		{
			name: "failure: request query paramters",
			fields: fields{
				Title:       "Schema Test",
				Description: "testing validation",
				Parameters: map[string]ParameterProperties{
					"param1": {
						Description: "String",
						Example:     "Test",
						Validation: ParameterValidation{
							String: &ParameterStringValidator{},
						},
					},
					"param2": {
						Description: "Number",
						Example:     "Test",
						Validation: ParameterValidation{
							Number: &ParameterDataNumberValidator{},
						},
					},
					"param3": {
						Description: "Boolean",
						Example:     "Test",
						Validation: ParameterValidation{
							Boolean: &ParameterBooleanValidator{},
						},
					},
					"param4": {
						Description:          "Number Array",
						Example:              "Test",
						InlineArray:          true,
						InlineArraySeperator: ",",
						Validation: ParameterValidation{
							Number: &ParameterDataNumberValidator{},
						},
					},
				},
			},
			args: args{
				req: func() *http.Request {
					r := httptest.NewRequest(http.MethodGet, "https://www.google.com", nil)
					q := r.URL.Query()
					q.Add("param1", "some string")
					q.Add("param2", "42")
					q.Add("param3", "true")
					q.Add("param4", "13,55")
					q.Add("param5", "13,55")
					r.URL.RawQuery = q.Encode()
					return r
				}(),
			},
			wantErr: true,
		},
		{
			name: "failure: wrong data type",
			fields: fields{
				Title:       "Schema Test",
				Description: "testing validation",
				Parameters: map[string]ParameterProperties{
					"param1": {
						Description: "String",
						Example:     "Test",
						Validation: ParameterValidation{
							String: &ParameterStringValidator{},
						},
					},
					"param2": {
						Description: "Number",
						Example:     "Test",
						Validation: ParameterValidation{
							Number: &ParameterDataNumberValidator{},
						},
					},
					"param3": {
						Description: "Boolean",
						Example:     "Test",
						Validation: ParameterValidation{
							Boolean: &ParameterBooleanValidator{},
						},
					},
					"param4": {
						Description:          "Number Array",
						Example:              "Test",
						InlineArray:          true,
						InlineArraySeperator: ",",
						Validation: ParameterValidation{
							Number: &ParameterDataNumberValidator{},
						},
					},
				},
			},
			args: args{
				req: func() *http.Request {
					r := httptest.NewRequest(http.MethodGet, "https://www.google.com", nil)
					q := r.URL.Query()
					q.Add("param1", "some string")
					q.Add("param2", "42")
					q.Add("param3", "string")
					q.Add("param4", "13,55")
					r.URL.RawQuery = q.Encode()
					return r
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
				Parameters:  tt.fields.Parameters,
			}
			if err := s.Validate(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("Schema.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
