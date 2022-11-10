package main

import (
	"reflect"
	"testing"
)

func Test_spiralStep(t *testing.T) {
	type args struct {
		matrix [][]int32
	}
	tests := []struct {
		name string
		args args
		want [][]int32
	}{
		{
			name: "spiral step test case 1 with 5 x 5 dimension",
			args: args{matrix: [][]int32{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
				{11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20},
				{21, 22, 23, 24, 25},
			}},
			want: [][]int32{
				{1, 2, 3, 4, 5},
				{10, 15, 20, 25},
				{24, 23, 22, 21},
				{16, 11, 6},
				{7, 8, 9},
				{14, 19},
				{18, 17},
				{12},
				{13},
			},
		},{
			name: "spiral step test case 2 with 1 x 1 dimension",
			args: args{matrix: [][]int32{
				{1},
			}},
			want: [][]int32{
				{1},
			},
		},{
			name: "spiral step test case 3 with 5 x 2 dimension",
			args: args{matrix: [][]int32{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
			}},
			want: [][]int32{
				{1, 2, 3, 4, 5},
				{10},
				{9, 8, 7, 6},
			},
		},{
			name: "spiral step test case 4 with nil value in parameter",
			args: args{matrix: nil},
			want: nil,
		},{
			name: "spiral step test case 5 with parameter length = 0",
			args: args{matrix: [][]int32{}},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spiralStep(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralStep() = %v, want %v", got, tt.want)
			}
		})
	}
}
