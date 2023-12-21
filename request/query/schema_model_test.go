package query

import (
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
					Title:           "Test",
					Description:     "test example",
					LooseValidation: false,
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
			name: "success: loose",
			args: args{
				schema: Schema{
					Title:           "Test",
					Description:     "test example",
					LooseValidation: true,
				},
			},
			wantErr: false,
		},
		{
			name: "fail: no title",
			args: args{
				schema: Schema{
					Description:     "test example",
					LooseValidation: false,
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
					Title:           "Test",
					Description:     "test example",
					LooseValidation: false,
				},
			},
			wantErr: true,
		},
		{
			name: "fail: parameter properties",
			args: args{
				schema: Schema{
					Title:           "Test",
					Description:     "test example",
					LooseValidation: false,
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
