package parameter

import (
	"testing"
	"time"
)

func TestParameterTimeValidation_Validate_Format(t *testing.T) {
	type fields struct {
		Format string
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
				Format: time.RFC3339,
			},
			args: args{
				value: "2023-10-12T07:20:50.52Z",
			},
			wantErr: false,
		},
		{
			name: "failure",
			fields: fields{
				Format: time.RFC822,
			},
			args: args{
				value: "2023-10-12T07:20:50.52Z",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := TimeValidator{
				Format: tt.fields.Format,
			}
			if err := p.Validate(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("ParameterTimeValidation.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParameterTimeValidation_Validate_Value(t *testing.T) {
	type fields struct {
		Format string
		Value  *string
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
			name: " success",
			fields: fields{
				Format: time.RFC3339,
				Value: func() *string {
					t := "2023-10-12T07:20:50.52Z"
					return &t
				}(),
			},
			args: args{
				value: "2023-10-12T07:20:50.52Z",
			},
			wantErr: false,
		},
		{
			name: " success",
			fields: fields{
				Format: time.RFC3339,
				Value: func() *string {
					t := "2043-10-12T07:20:50.52Z"
					return &t
				}(),
			},
			args: args{
				value: "2023-10-12T07:20:50.52Z",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := TimeValidator{
				Format: tt.fields.Format,
				Value:  tt.fields.Value,
			}
			if err := p.Validate(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("ParameterTimeValidation.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParameterTimeValidation_Validate_Before(t *testing.T) {
	type fields struct {
		Format string
		Before *string
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
			name: " success",
			fields: fields{
				Format: time.RFC3339,
				Before: func() *string {
					t := "2024-10-12T07:20:50.52Z"
					return &t
				}(),
			},
			args: args{
				value: "2023-10-12T07:20:50.52Z",
			},
			wantErr: false,
		},
		{
			name: " failure",
			fields: fields{
				Format: time.RFC3339,
				Before: func() *string {
					t := "2022-10-12T07:20:50.52Z"
					return &t
				}(),
			},
			args: args{
				value: "2023-10-12T07:20:50.52Z",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := TimeValidator{
				Format: tt.fields.Format,
				Before: tt.fields.Before,
			}
			if err := p.Validate(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("ParameterTimeValidation.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParameterTimeValidation_Validate_After(t *testing.T) {
	type fields struct {
		Format string
		After  *string
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
			name: " success",
			fields: fields{
				Format: time.RFC3339,
				After: func() *string {
					t := "2022-10-12T07:20:50.52Z"
					return &t
				}(),
			},
			args: args{
				value: "2023-10-12T07:20:50.52Z",
			},
			wantErr: false,
		},
		{
			name: " failure",
			fields: fields{
				Format: time.RFC3339,
				After: func() *string {
					t := "2024-10-12T07:20:50.52Z"
					return &t
				}(),
			},
			args: args{
				value: "2023-10-12T07:20:50.52Z",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := TimeValidator{
				Format: tt.fields.Format,
				After:  tt.fields.After,
			}
			if err := p.Validate(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("ParameterTimeValidation.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTimeArrayValidator_Validate_Values(t *testing.T) {
	type fields struct {
		Format string
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
			name: " success",
			fields: fields{
				Format: time.RFC3339,
				Values: []string{"2023-10-12T07:20:50.52Z", "2023-07-12T07:20:50.52Z"},
			},
			args: args{
				values: []string{"2023-10-12T07:20:50.52Z", "2023-07-12T07:20:50.52Z"},
			},
			wantErr: false,
		},
		{
			name: " failure",
			fields: fields{
				Format: time.RFC3339,
				Values: []string{"2023-10-12T07:20:50.52Z", "2023-07-12T07:20:50.52Z"},
			},
			args: args{
				values: []string{"2023-10-12T07:20:50.52Z", "2022-07-12T07:20:50.52Z"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tav := TimeArrayValidator{
				Format: tt.fields.Format,
				Values: tt.fields.Values,
			}
			if err := tav.Validate(tt.args.values); (err != nil) != tt.wantErr {
				t.Errorf("TimeArrayValidator.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTimeArrayValidator_Validate_Before_After(t *testing.T) {
	type fields struct {
		Format string
		Before *string
		After  *string
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
			name: " success",
			fields: fields{
				Format: time.RFC3339,
				Before: func() *string {
					t := "2025-10-12T07:20:50.52Z"
					return &t
				}(),
				After: func() *string {
					t := "2020-10-12T07:20:50.52Z"
					return &t
				}(),
			},
			args: args{
				values: []string{"2023-10-12T07:20:50.52Z", "2023-07-12T07:20:50.52Z"},
			},
			wantErr: false,
		},
		{
			name: "failure: before",
			fields: fields{
				Format: time.RFC3339,
				Before: func() *string {
					t := "2021-10-12T07:20:50.52Z"
					return &t
				}(),
				After: func() *string {
					t := "2020-10-12T07:20:50.52Z"
					return &t
				}(),
			},
			args: args{
				values: []string{"2023-10-12T07:20:50.52Z", "2023-07-12T07:20:50.52Z"},
			},
			wantErr: true,
		},
		{
			name: "failure: after",
			fields: fields{
				Format: time.RFC3339,
				Before: func() *string {
					t := "2025-10-12T07:20:50.52Z"
					return &t
				}(),
				After: func() *string {
					t := "2024-10-12T07:20:50.52Z"
					return &t
				}(),
			},
			args: args{
				values: []string{"2023-10-12T07:20:50.52Z", "2023-07-12T07:20:50.52Z"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tav := TimeArrayValidator{
				Format: tt.fields.Format,
				Before: tt.fields.Before,
				After:  tt.fields.After,
			}
			if err := tav.Validate(tt.args.values); (err != nil) != tt.wantErr {
				t.Errorf("TimeArrayValidator.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
