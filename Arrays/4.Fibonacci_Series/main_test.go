package main

import (
	"reflect"
	"testing"
)

func Test_fibonacci(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Negative Number",
			args: args{-2},
			want: nil,
		}, {
			name: "Zero Number",
			args: args{0},
			want: nil,
		}, {
			name: "Positive Number - Test 1",
			args: args{10},
			want: []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34},
		}, {
			name: "Positive Number - Test 2",
			args: args{15},
			want: []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibonacci(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
