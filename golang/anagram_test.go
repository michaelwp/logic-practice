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
				s1: "Danger",
				s2: "Garden",
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

func Test_sherlockAndAnagrams(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "test 1",
			args: args{s: "ifailuhkqq"},
			want: 3,
		},{
			name: "test 2",
			args: args{s: "kkkk"},
			want: 10,
		},{
			name: "test 3",
			args: args{s: "mom"},
			want: 2,
		},{
			name: "test 4",
			args: args{s: "abba"},
			want: 4,
		},{
			name: "test 5",
			args: args{s: "abcd"},
			want: 0,
		},{
			name: "test 6",
			args: args{s: "cdcd"},
			want: 5,
		},{
			name: "test 7",
			args: args{s: ""},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sherlockAndAnagrams(tt.args.s); got != tt.want {
				t.Errorf("sherlockAndAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
