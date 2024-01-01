package field

import "testing"

func TestRequired_Validate(t *testing.T) {
	type fields struct {
		OneOf   [][]string
		Present map[string][]string
	}
	type args struct {
		fields map[string]struct{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success: one of",
			fields: fields{
				OneOf: [][]string{{"field1"}, {"fieldA"}},
			},
			args: args{
				fields: map[string]struct{}{
					"field1": {},
					"field2": {},
					"field3": {},
					"field4": {},
					"field5": {},
					"field6": {},
					"field7": {},
				},
			},
			wantErr: false,
		},
		{
			name: "success: present",
			fields: fields{
				Present: map[string][]string{
					"field2": {"field5", "field4"},
				},
			},
			args: args{
				fields: map[string]struct{}{
					"field1": {},
					"field2": {},
					"field3": {},
					"field4": {},
					"field5": {},
					"field6": {},
					"field7": {},
				},
			},
			wantErr: false,
		},
		{
			name: "success: if one is present, then others are required",
			fields: fields{
				OneOf: [][]string{{"field1"}, {"fieldA"}},
				Present: map[string][]string{
					"field1": {"field5", "field4"},
				},
			},
			args: args{
				fields: map[string]struct{}{
					"field1": {},
					"field2": {},
					"field3": {},
					"field4": {},
					"field5": {},
					"field6": {},
					"field7": {},
				},
			},
			wantErr: false,
		},
		{
			name: "failure: one of",
			fields: fields{
				OneOf: [][]string{{"fieldB"}, {"fieldA"}},
			},
			args: args{
				fields: map[string]struct{}{
					"field1": {},
					"field2": {},
					"field3": {},
					"field4": {},
					"field5": {},
					"field6": {},
					"field7": {},
				},
			},
			wantErr: true,
		},
		{
			name: "failure: present",
			fields: fields{
				Present: map[string][]string{
					"field2": {"fieldG", "fieldX"},
				},
			},
			args: args{
				fields: map[string]struct{}{
					"field1": {},
					"field2": {},
					"field3": {},
					"field4": {},
					"field5": {},
					"field6": {},
					"field7": {},
				},
			},
			wantErr: true,
		},
		{
			name: "failure: if one is present, then others are required",
			fields: fields{
				OneOf: [][]string{{"field1"}, {"fieldA"}},
				Present: map[string][]string{
					"field1": {"fieldF", "fieldK"},
				},
			},
			args: args{
				fields: map[string]struct{}{
					"field1": {},
					"field2": {},
					"field3": {},
					"field4": {},
					"field5": {},
					"field6": {},
					"field7": {},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Required{
				OneOf:   tt.fields.OneOf,
				Present: tt.fields.Present,
			}
			if err := r.Validate(tt.args.fields); (err != nil) != tt.wantErr {
				t.Errorf("Required.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
