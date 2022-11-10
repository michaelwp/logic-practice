package main

import "testing"

func Test_minCostPath(t *testing.T) {
	type args struct {
		cost [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{cost: [][]int{
				{3, 2, 12, 15, 10},
				{6, 19, 7, 11, 17},
				{8, 5, 12, 32, 21},
				{3, 20, 2, 9, 7},
			}},
			want: 32,
		},
		{
			name: "test 2",
			args: args{cost: [][]int{
				{1, 2, 3},
				{4, 8, 2},
				{1, 5, 3},
			}},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostPath(tt.args.cost); got != tt.want {
				t.Errorf("minCostPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
