package main

import "testing"

func Test_isAnagram(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test case 1",
			args: args{
				s1: "word",
				s2: "drow",
			},
			want: true,
		},{
			name: "test case 2",
			args: args{
				s1: "superman",
				s2: "superwoman",
			},
			want: false,
		},{
			name: "test case 3",
			args: args{
				s1: "danger",
				s2: "garden",
			},
			want: true,
		},{
			name: "test case 4",
			args: args{
				s1: "finaccel",
				s2: "finacell",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("isAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
