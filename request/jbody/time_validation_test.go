package jbody

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/g8rswimmer/httpx/request/internal/parameter"
)

func TestTimeValidator_Validate(t *testing.T) {
	type fields struct {
		TimeValidator parameter.TimeValidator
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
				TimeValidator: parameter.TimeValidator{
					Format: time.RFC3339,
					Value: func() *string {
						t := "2019-10-12T07:20:50.52Z"
						return &t
					}(),
				},
			},
			args: args{
				value: func() any {
					str := `{"value":"2019-10-12T07:20:50.52Z"}`
					var obj map[string]any
					_ = json.Unmarshal([]byte(str), &obj)
					return obj["value"]
				}(),
			},
			wantErr: false,
		},
		{
			name: "success",
			fields: fields{
				TimeValidator: parameter.TimeValidator{
					Format: time.RFC3339,
					Value: func() *string {
						t := "2019-10-12T07:20:50.52Z"
						return &t
					}(),
				},
			},
			args: args{
				value: func() any {
					str := `{"value":2019}`
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
			tr := TimeValidator{
				TimeValidator: tt.fields.TimeValidator,
			}
			if err := tr.Validate(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("TimeValidator.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTimeArrayValidator_Validate(t *testing.T) {
	type fields struct {
		TimeArrayValidator parameter.TimeArrayValidator
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
				TimeArrayValidator: parameter.TimeArrayValidator{
					Format: time.RFC3339,
					Values: []string{"2019-10-12T07:20:50.52Z", "2024-12-12T07:20:50.52Z"},
				},
			},
			args: args{
				value: func() any {
					str := `{"value":["2019-10-12T07:20:50.52Z","2024-12-12T07:20:50.52Z"]}`
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
				TimeArrayValidator: parameter.TimeArrayValidator{
					Format: time.RFC3339,
					Values: []string{"2019-10-12T07:20:50.52Z", "2024-12-12T07:20:50.52Z"},
				},
			},
			args: args{
				value: func() any {
					str := `{"value":[55.25]}`
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
			s := TimeArrayValidator{
				TimeArrayValidator: tt.fields.TimeArrayValidator,
			}
			if err := s.Validate(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("TimeArrayValidator.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
