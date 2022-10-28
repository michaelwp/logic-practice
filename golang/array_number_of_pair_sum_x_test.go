package main

import (
	"reflect"
	"testing"
)

func Test_numberOfPairSumX(t *testing.T) {
	type args struct {
		arr []int
		sum int
	}
	tests := []struct {
		name    string
		args    args
		wantRes []string
	}{
		{
			name: "test1",
			args: args{
				arr: []int{1, 5, 3, 2, 4},
				sum: 5,
			},
			wantRes: []string{
				"{1,4}", "{2,3}",
			},
		},
		{
			name: "test2",
			args: args{
				arr: []int{1, 5, 3, 2, 4},
				sum: 2,
			},
			wantRes: []string{},
		},
		{
			name: "test3",
			args: args{
				arr: []int{10, 50, 5, 7, -35, 8},
				sum: 15,
			},
			wantRes: []string{"{-35,50}", "{5,10}", "{7,8}"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := numberOfPairSumX(tt.args.arr, tt.args.sum); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("numberOfPairSumX() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
