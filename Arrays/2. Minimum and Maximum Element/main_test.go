package main

import (
	"reflect"
	"testing"
)

func Test_findMinMax(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 int
		want2 int
	}{
		{
			name: "Empty Array",
			args: args{
				[]int{},
			},
			want:  true,
			want1: 0,
			want2: 0,
		},
		{
			name: "Single Element In Array",
			args: args{
				[]int{67},
			},
			want:  false,
			want1: 67,
			want2: 67,
		},
		{
			name: "Multiple Elements In Array",
			args: args{
				[]int{101, 233, 43, 56, 0, -21, 90, -49},
			},
			want:  false,
			want1: -49,
			want2: 233,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := findMinMax(tt.args.arr)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMinMax() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("findMinMax() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("findMinMax() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
