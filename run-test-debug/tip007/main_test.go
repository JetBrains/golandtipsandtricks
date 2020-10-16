package main

import (
	"fmt"
	"testing"
)

// Step 1. Run the test below and observe the subtests
//         getting new buttons on the gutter

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
		{"1", fields{1}, args{0}, 1},
		{"2", fields{1}, args{1}, 2},
		{"3", fields{1}, args{1}, 3},
		{"4", fields{2}, args{2}, 4},
		{"5", fields{3}, args{2}, 4},
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

func Test_demoType_AddToField_withParams(t *testing.T) {
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
		{"1", fields{1}, args{0}, 1},
		{"2", fields{1}, args{1}, 2},
		{"3", fields{1}, args{1}, 3},
		{"4", fields{2}, args{2}, 4},
		{"5", fields{3}, args{2}, 4},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("test num %s", tt.name), func(t *testing.T) {
			d := &demoType{
				Field1: tt.fields.Field1,
			}
			if got := d.AddToField(tt.args.val); got != tt.want {
				t.Errorf("AddToField() = %v, want %v", got, tt.want)
			}
		})
	}
}
