package jbody

import (
	"encoding/json"
	"testing"

	"github.com/g8rswimmer/httpx/request/internal/parameter"
)

func TestBooleanValidator_Validate(t *testing.T) {
	type fields struct {
		BooleanValidator parameter.BooleanValidator
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
				BooleanValidator: parameter.BooleanValidator{
					Value: func() *bool {
						b := true
						return &b
					}(),
				},
			},
			args: args{
				value: func() any {
					str := `{"value":true}`
					var obj map[string]any
					_ = json.Unmarshal([]byte(str), &obj)
					return obj["value"]
				}(),
			},
			wantErr: false,
		},
		{
			name: "failure",
			fields: fields{
				BooleanValidator: parameter.BooleanValidator{
					Value: func() *bool {
						b := true
						return &b
					}(),
				},
			},
			args: args{
				value: func() any {
					str := `{"value":56}`
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
			b := BooleanValidator{
				BooleanValidator: tt.fields.BooleanValidator,
			}
			if err := b.Validate(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("BooleanValidator.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
