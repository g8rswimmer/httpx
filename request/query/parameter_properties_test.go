package query

import (
	"testing"
	"time"

	"github.com/g8rswimmer/httpx/request/internal/parameter"
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
			name: "success",
			args: args{
				properties: ParameterProperties{
					Description: "None",
					Example:     "Test",
					Validation: ParameterValidation{
						String: &parameter.StringValidator{},
					},
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
					InlineArray: true,
					Validation: ParameterValidation{
						String: &parameter.StringValidator{},
					},
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

func TestParameterProperties_Validate(t *testing.T) {
	type fields struct {
		Description          string
		Example              string
		DataType             string
		InlineArray          bool
		InlineArraySeperator string
		Validation           ParameterValidation
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
				Validation: ParameterValidation{
					String: &parameter.StringValidator{},
				},
			},
			args: args{
				value: "some string",
			},
			wantErr: false,
		},
		{
			name: "success: string array",
			fields: fields{
				Description:          "string test",
				Example:              "none",
				InlineArray:          true,
				InlineArraySeperator: ",",
				Validation: ParameterValidation{
					StringArray: &parameter.StringArrayValidator{},
				},
			},
			args: args{
				value: "some string,hello",
			},
			wantErr: false,
		},
		{
			name: "success: number",
			fields: fields{
				Description: "number test",
				Example:     "none",
				Validation: ParameterValidation{
					Number: &NumberValidator{},
				},
			},
			args: args{
				value: "42",
			},
			wantErr: false,
		},
		{
			name: "success: number array",
			fields: fields{
				Description:          "number test",
				Example:              "none",
				InlineArray:          true,
				InlineArraySeperator: ",",
				Validation: ParameterValidation{
					NumberArray: &NumberArrayValidator{},
				},
			},
			args: args{
				value: "42,88",
			},
			wantErr: false,
		},
		{
			name: "success: boolean",
			fields: fields{
				Description: "boolean test",
				Example:     "none",
				Validation: ParameterValidation{
					Boolean: &BooleanValidator{},
				},
			},
			args: args{
				value: "true",
			},
			wantErr: false,
		},
		{
			name: "success: time",
			fields: fields{
				Description: "time test",
				Example:     "none",
				Validation: ParameterValidation{
					Time: &parameter.TimeValidator{
						Format: time.RFC3339,
					},
				},
			},
			args: args{
				value: "2023-10-12T07:20:50.52Z",
			},
			wantErr: false,
		},
		{
			name: "success: time",
			fields: fields{
				Description:          "time test",
				Example:              "none",
				InlineArray:          true,
				InlineArraySeperator: ",",
				Validation: ParameterValidation{
					TimeArray: &parameter.TimeArrayValidator{
						Format: time.RFC3339,
					},
				},
			},
			args: args{
				value: "2023-10-12T07:20:50.52Z,2025-10-12T07:20:50.52Z",
			},
			wantErr: false,
		},
		{
			name: "success: optional with value",
			fields: fields{
				Description: "optional string test",
				Example:     "none",
				Validation: ParameterValidation{
					String: &parameter.StringValidator{},
				},
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
				Validation: ParameterValidation{
					String: &parameter.StringValidator{},
				},
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
				InlineArray:          true,
				InlineArraySeperator: ",",
				Validation: ParameterValidation{
					NumberArray: &NumberArrayValidator{},
				},
			},
			args: args{
				value: "42,78",
			},
			wantErr: false,
		},
		{
			name: "failure: number",
			fields: fields{
				Description: "number test",
				Example:     "none",
				Validation: ParameterValidation{
					Number: &NumberValidator{},
				},
			},
			args: args{
				value: "42xxx",
			},
			wantErr: true,
		},
		{
			name: "failure: number array",
			fields: fields{
				Description:          "number test",
				Example:              "none",
				InlineArray:          true,
				InlineArraySeperator: ",",
				Validation: ParameterValidation{
					NumberArray: &NumberArrayValidator{},
				},
			},
			args: args{
				value: "42xxx,55",
			},
			wantErr: true,
		},
		{
			name: "failure: boolean",
			fields: fields{
				Description: "boolean test",
				Example:     "none",
				Validation: ParameterValidation{
					Boolean: &BooleanValidator{},
				},
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
				InlineArray:          tt.fields.InlineArray,
				InlineArraySeperator: tt.fields.InlineArraySeperator,
				Validation:           tt.fields.Validation,
			}
			if err := p.Validate(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("ParameterProperties.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
