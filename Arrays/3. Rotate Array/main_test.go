package main

import (
	"reflect"
	"testing"
)

func Test_rotate(t *testing.T) {
	type args struct {
		arr []int
		n   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Empty Array",
			args: args{
				arr: []int{},
				n:   2,
			},
			want: []int{},
		},
		{
			name: "Single Element In Array With Negative Rotations",
			args: args{
				arr: []int{45},
				n:   -3,
			},
			want: []int{45},
		},
		{
			name: "Single Element In Array With Zero Rotations",
			args: args{
				arr: []int{45},
				n:   0,
			},
			want: []int{45},
		},
		{
			name: "Single Element In Array With Some Positive Number Rotations",
			args: args{
				arr: []int{45},
				n:   1,
			},
			want: []int{45},
		},
		{
			name: "Multiple Element In Array With Number of Rotations Equal to Array Length",
			args: args{
				arr: []int{10, 20, 30, 40, 50},
				n:   5,
			},
			want: []int{10, 20, 30, 40, 50},
		},
		{
			name: "Multiple Element In Array With Postive Rotations - Test 1",
			args: args{
				arr: []int{10, 20, 30, 40, 50},
				n:   4,
			},
			want: []int{20, 30, 40, 50, 10},
		},
		{
			name: "Multiple Element In Array With Postive Rotations - Test 2",
			args: args{
				arr: []int{10, 20, 30, 40, 50},
				n:   2,
			},
			want: []int{40, 50, 10, 20, 30},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotate(tt.args.arr, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotate() = %v, want %v", got, tt.want)
			}
		})
	}
}
