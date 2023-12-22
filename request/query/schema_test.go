package query

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSchemaModelParameterPropertiesValidator(t *testing.T) {
	type args struct {
		properties ParameterProperties
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success: string",
			args: args{
				properties: ParameterProperties{
					Description: "None",
					Example:     "Test",
					DataType:    DataTypeString,
				},
			},
			wantErr: false,
		},
		{
			name: "success: number",
			args: args{
				properties: ParameterProperties{
					Description: "None",
					Example:     "Test",
					DataType:    DataTypeNumber,
				},
			},
			wantErr: false,
		},
		{
			name: "success: boolean",
			args: args{
				properties: ParameterProperties{
					Description: "None",
					Example:     "Test",
					DataType:    DataTypeBoolean,
				},
			},
			wantErr: false,
		},
		{
			name: "success: inline string array",
			args: args{
				properties: ParameterProperties{
					Description:          "None",
					Example:              "Test",
					DataType:             DataTypeString,
					InlineArray:          true,
					InlineArraySeperator: ",",
				},
			},
			wantErr: false,
		},
		{
			name: "fail: bad data type",
			args: args{
				properties: ParameterProperties{
					Description: "None",
					Example:     "Test",
					DataType:    "not found",
				},
			},
			wantErr: true,
		},
		{
			name: "fail: no inline array seperator",
			args: args{
				properties: ParameterProperties{
					Description: "None",
					Example:     "Test",
					DataType:    DataTypeString,
					InlineArray: true,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SchemaModelParameterPropertiesValidator(tt.args.properties); (err != nil) != tt.wantErr {
				t.Errorf("SchemaModelParameterPropertiesValidator() error = %v, wantErr %v", err, tt.wantErr)
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
					Title:       "Test",
					Description: "test example",
					Parameters: map[string]ParameterProperties{
						"param1": {
							Description: "None",
							Example:     "Test",
							DataType:    DataTypeString,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "fail: no title",
			args: args{
				schema: Schema{
					Description: "test example",
					Parameters: map[string]ParameterProperties{
						"param1": {
							Description: "None",
							Example:     "Test",
							DataType:    DataTypeString,
						},
					},
				},
			},
			wantErr: true,
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
							DataType:    "none",
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

func TestParameterProperties_Validate(t *testing.T) {
	type fields struct {
		Description          string
		Example              string
		DataType             string
		InlineArray          bool
		InlineArraySeperator string
		Optional             bool
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
				Description: "string test",
				Example:     "none",
				DataType:    DataTypeString,
			},
			args: args{
				value: "some string",
			},
			wantErr: false,
		},
		{
			name: "success: number",
			fields: fields{
				Description: "number test",
				Example:     "none",
				DataType:    DataTypeNumber,
			},
			args: args{
				value: "42",
			},
			wantErr: false,
		},
		{
			name: "success: boolean",
			fields: fields{
				Description: "boolean test",
				Example:     "none",
				DataType:    DataTypeBoolean,
			},
			args: args{
				value: "true",
			},
			wantErr: false,
		},
		{
			name: "success: optional with value",
			fields: fields{
				Description: "optional string test",
				Example:     "none",
				DataType:    DataTypeString,
				Optional:    true,
			},
			args: args{
				value: "some string",
			},
			wantErr: false,
		},
		{
			name: "success: optional with no value",
			fields: fields{
				Description: "optional string test",
				Example:     "none",
				DataType:    DataTypeString,
				Optional:    true,
			},
			args: args{
				value: "",
			},
			wantErr: false,
		},
		{
			name: "success: inline array number",
			fields: fields{
				Description:          "number array test",
				Example:              "none",
				DataType:             DataTypeNumber,
				InlineArray:          true,
				InlineArraySeperator: ",",
			},
			args: args{
				value: "42,78",
			},
			wantErr: false,
		},
		{
			name: "failure: required value",
			fields: fields{
				Description: "string test",
				Example:     "none",
				DataType:    DataTypeString,
			},
			args: args{
				value: "",
			},
			wantErr: true,
		},
		{
			name: "failure: number",
			fields: fields{
				Description: "number test",
				Example:     "none",
				DataType:    DataTypeNumber,
			},
			args: args{
				value: "42xxx",
			},
			wantErr: true,
		},
		{
			name: "failure: boolean",
			fields: fields{
				Description: "boolean test",
				Example:     "none",
				DataType:    DataTypeBoolean,
			},
			args: args{
				value: "not",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ParameterProperties{
				Description:          tt.fields.Description,
				Example:              tt.fields.Example,
				DataType:             tt.fields.DataType,
				InlineArray:          tt.fields.InlineArray,
				InlineArraySeperator: tt.fields.InlineArraySeperator,
				Optional:             tt.fields.Optional,
			}
			if err := p.Validate(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("ParameterProperties.Validate() error = %v, wantErr %v", err, tt.wantErr)
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
						DataType:    DataTypeString,
					},
					"param2": {
						Description: "Number",
						Example:     "Test",
						DataType:    DataTypeNumber,
					},
					"param3": {
						Description: "Boolean",
						Example:     "Test",
						DataType:    DataTypeBoolean,
					},
					"param4": {
						Description:          "Number Array",
						Example:              "Test",
						DataType:             DataTypeNumber,
						InlineArray:          true,
						InlineArraySeperator: ",",
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
						DataType:    DataTypeString,
					},
					"param2": {
						Description: "Number",
						Example:     "Test",
						DataType:    DataTypeNumber,
					},
					"param3": {
						Description: "Boolean",
						Example:     "Test",
						DataType:    DataTypeBoolean,
						Optional:    true,
					},
					"param4": {
						Description:          "Number Array",
						Example:              "Test",
						DataType:             DataTypeNumber,
						InlineArray:          true,
						InlineArraySeperator: ",",
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
						DataType:    DataTypeString,
					},
					"param2": {
						Description: "Number",
						Example:     "Test",
						DataType:    DataTypeNumber,
					},
					"param3": {
						Description: "Boolean",
						Example:     "Test",
						DataType:    DataTypeBoolean,
					},
					"param4": {
						Description:          "Number Array",
						Example:              "Test",
						DataType:             DataTypeNumber,
						InlineArray:          true,
						InlineArraySeperator: ",",
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
						DataType:    DataTypeString,
					},
					"param2": {
						Description: "Number",
						Example:     "Test",
						DataType:    DataTypeNumber,
					},
					"param3": {
						Description: "Boolean",
						Example:     "Test",
						DataType:    DataTypeBoolean,
					},
					"param4": {
						Description:          "Number Array",
						Example:              "Test",
						DataType:             DataTypeNumber,
						InlineArray:          true,
						InlineArraySeperator: ",",
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