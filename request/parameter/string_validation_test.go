package parameter

import (
	"testing"
)

func TestParameterStringValidation_Validate_Value(t *testing.T) {
	type fields struct {
		Value *string
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
			name: "success",
			fields: fields{
				Value: func() *string {
					s := "test"
					return &s
				}(),
			},
			args: args{
				value: "test",
			},
			wantErr: false,
		},
		{
			name: "failure",
			fields: fields{
				Value: func() *string {
					s := "test"
					return &s
				}(),
			},
			args: args{
				value: "fail",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := StringValidator{
				Value: tt.fields.Value,
			}
			if err := p.Validate(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("ParameterStringValidation.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParameterStringValidation_Validate_OneOf(t *testing.T) {
	type fields struct {
		OneOf []string
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
			name: "success",
			fields: fields{
				OneOf: []string{"one", "two"},
			},
			args: args{
				value: "two",
			},
			wantErr: false,
		},
		{
			name: "failure",
			fields: fields{
				OneOf: []string{"one", "two"},
			},
			args: args{
				value: "three",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := StringValidator{
				OneOf: tt.fields.OneOf,
			}
			if err := p.Validate(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("ParameterStringValidation.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParameterStringValidation_Validate_RegEx(t *testing.T) {
	type fields struct {
		RegEx *string
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
			name: "success: uuidv4",
			fields: fields{
				RegEx: func() *string {
					s := RegExUUIDv4
					return &s
				}(),
			},
			args: args{
				value: "ead5b1f7-e8e3-4c65-ba1d-29c98ea5a5b8",
			},
			wantErr: false,
		},
		{
			name: "failure: uuidv4",
			fields: fields{
				RegEx: func() *string {
					s := RegExUUIDv4
					return &s
				}(),
			},
			args: args{
				value: "eadb1f7-e8e3-4c65-ba1d-29c98ea5a5b8",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := StringValidator{
				RegEx: tt.fields.RegEx,
			}
			if err := p.Validate(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("ParameterStringValidation.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStringArrayValidator_Validate_Value(t *testing.T) {
	type fields struct {
		Values []string
	}
	type args struct {
		values []string
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
				Values: []string{"one", "two", "three"},
			},
			args: args{
				values: []string{"one", "two", "three"},
			},
			wantErr: false,
		},
		{
			name: "failure: match",
			fields: fields{
				Values: []string{"one", "two", "four"},
			},
			args: args{
				values: []string{"one", "two", "three"},
			},
			wantErr: true,
		},
		{
			name: "failure: length",
			fields: fields{
				Values: []string{"one", "two", "three"},
			},
			args: args{
				values: []string{"one", "two", "three", "four"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StringArrayValidator{
				Values: tt.fields.Values,
			}
			if err := s.Validate(tt.args.values); (err != nil) != tt.wantErr {
				t.Errorf("StringArrayValidator.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStringArrayValidator_Validate_RegEx(t *testing.T) {
	type fields struct {
		RegEx *string
	}
	type args struct {
		values []string
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
				RegEx: func() *string {
					s := RegExUUIDv4
					return &s
				}(),
			},
			args: args{
				values: []string{"b22df31f-214b-4a22-9a24-f95909fbfafa", "8b4a60d8-c203-460f-92eb-82646c93d792"},
			},
			wantErr: false,
		},
		{
			name: "failure",
			fields: fields{
				RegEx: func() *string {
					s := RegExUUIDv4
					return &s
				}(),
			},
			args: args{
				values: []string{"b22df31f-214b-4a22-9a24-f95909fbfafa", "8b4a60d8-c203-460f-92eb-82646c93d792", "nope"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StringArrayValidator{
				RegEx: tt.fields.RegEx,
			}
			if err := s.Validate(tt.args.values); (err != nil) != tt.wantErr {
				t.Errorf("StringArrayValidator.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStringArrayValidator_Validate_Present(t *testing.T) {
	type fields struct {
		Present []string
	}
	type args struct {
		values []string
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
				Present: []string{"one", "three"},
			},
			args: args{
				values: []string{"one", "two", "three"},
			},
			wantErr: false,
		},
		{
			name: "failure",
			fields: fields{
				Present: []string{"one", "four"},
			},
			args: args{
				values: []string{"one", "two", "three"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StringArrayValidator{
				Present: tt.fields.Present,
			}
			if err := s.Validate(tt.args.values); (err != nil) != tt.wantErr {
				t.Errorf("StringArrayValidator.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
