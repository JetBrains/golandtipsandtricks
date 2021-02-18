package main

import "testing"

func Test_demoType_AddToField(t *testing.T) {
	type fields struct {
		Field1 int
	}
	type args struct {
		val int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "first",
			fields: fields{Field1: 1},
			args:   args{val: 1},
			want:   2,
		},
		{
			name:   "second",
			fields: fields{Field1: 1},
			args:   args{val: 1},
			want:   2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &demoType{
				Field1: tt.fields.Field1,
			}
			if got := d.AddToField(tt.args.val); got != tt.want {
				t.Errorf("AddToField() = %v, want %v", got, tt.want)
			}
		})
	}
}
