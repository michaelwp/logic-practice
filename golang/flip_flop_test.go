package main

import (
	"reflect"
	"testing"
)

func Test_flipFlop(t *testing.T) {
	type args struct {
		number int
		key1   int
		key2   int
	}
	tests := []struct {
		name    string
		args    args
		wantRes map[int]string
	}{
		{
			name: "45",
			args: args{
				number: 45,
				key1:   3,
				key2:   5,
			},
			wantRes: map[int]string{
				3: "flip", 5: "flop", 6: "flip", 9: "flip", 10: "flop",
				12: "flip", 15: "flip flop", 18: "flip", 20: "flop", 21: "flip",
				24: "flip", 25: "flop", 27: "flip", 30: "flip flop", 33: "flip",
				35: "flop", 36: "flip", 39: "flip", 40: "flop", 42: "flip",
				45: "flip flop",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := flipFlop(tt.args.number, tt.args.key1, tt.args.key2); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("flipFlop() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
