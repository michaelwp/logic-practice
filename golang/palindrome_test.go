package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test 1",
			args: args{s: "0p"},
			want: false,
		},
		{
			name: "test 2",
			args: args{s: "A man, a plan, a canal: Panama"},
			want: true,
		},
		{
			name: "test 3",
			args: args{s: "race a car"},
			want: false,
		},
		{
			name: "test 4",
			args: args{s: " "},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
