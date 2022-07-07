package main

import (
	"reflect"
	"testing"
)

func TestArrayRotate(t *testing.T) {
	type args struct {
		arr []interface{}
		rotate int
	}

	type test struct {
		name string
		args args
		expectResp []interface{}
	}

	tests := []test{
		{
			name: "test 1",
			args: args{
				arr:    []interface{}{1, 2, 3, 4, 5},
				rotate: 3,
			},
			expectResp: []interface{}{3, 4, 5, 1, 2},
		},
		{
			name: "test 2",
			args: args{
				arr:    []interface{}{1, 2, 3, 4, 5},
				rotate: -2,
			},
			expectResp: []interface{}{3, 4, 5, 1, 2},
		},
		{
			name: "test 3",
			args: args{
				arr:    []interface{}{1, 2, 3, 4, 5},
				rotate: -11,
			},
			expectResp: []interface{}{2, 3, 4, 5, 1},
		},
		{
			name: "test 4",
			args: args{
				arr:    []interface{}{1, 2, 3, 4, 5},
				rotate: 17,
			},
			expectResp: []interface{}{4, 5, 1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualResp := rotate(tt.args.arr, tt.args.rotate)
			if !reflect.DeepEqual(actualResp, tt.expectResp) {
				t.Errorf("expect response: %v, \nactual response: %v", tt.expectResp, actualResp)
			}
		})
	}
}
