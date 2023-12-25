package parameter

import "testing"

func TestParameterBooleanValidation_Validate_Value(t *testing.T) {
	type fields struct {
		Value *bool
	}
	type args struct {
		value bool
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
				Value: func() *bool {
					b := true
					return &b
				}(),
			},
			args: args{
				value: true,
			},
			wantErr: false,
		},
		{
			name: "failure",
			fields: fields{
				Value: func() *bool {
					b := true
					return &b
				}(),
			},
			args: args{
				value: false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := BooleanValidator{
				Value: tt.fields.Value,
			}
			if err := p.Validate(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("ParameterBooleanValidation.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
