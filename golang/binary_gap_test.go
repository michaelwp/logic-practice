package main

import "testing"

func Test_binaryGap(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "tes 1",
			args: args{N: 32},
			want: 0,

		},
		{
			name: "tes 2",
			args: args{N: 15},
			want: 0,

		},
		{
			name: "tes 3",
			args: args{N: 1041},
			want: 5,

		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binaryGap(tt.args.N); got != tt.want {
				t.Errorf("binaryGap() = %v, want %v", got, tt.want)
			}
		})
	}
}
