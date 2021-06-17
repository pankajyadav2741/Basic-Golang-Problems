package main

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Zero Length Slice",
			args: args{
				[]int{},
			},
			want: []int{},
		},
		{
			name: "Single Element Length Slice",
			args: args{
				[]int{45},
			},
			want: []int{45},
		},
		{
			name: "Multiple Element Old-Length Slice",
			args: args{
				[]int{2, 0, 7, 9, -1, 15, 3},
			},
			want: []int{3, 15, -1, 9, 7, 0, 2},
		},
		{
			name: "Multiple Element Even-Length Slice",
			args: args{
				[]int{0, 2, 4, 6, 8, 10},
			},
			want: []int{10, 8, 6, 4, 2, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
