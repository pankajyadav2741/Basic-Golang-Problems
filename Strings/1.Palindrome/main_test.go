package main

import "testing"

func Test_palindrome(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Empty String",
			args: args{""},
			want: true,
		},
		{
			name: "Not Palindrome String",
			args: args{"hello"},
			want: false,
		},
		{
			name: "Palindrome String",
			args: args{"madam"},
			want: true,
		},
		{
			name: "Multi-String Not Palindrome",
			args: args{"hello world"},
			want: false,
		},
		{
			name: "Multi-String Palindrome",
			args: args{"madam madam"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := palindrome(tt.args.str); got != tt.want {
				t.Errorf("palindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
