package main

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestCommonDivisors(t *testing.T) {
	type args struct {
		x []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "divisors is 0",
			args: args{x: []int{320, 0}},
			want: nil,
		},
		{
			name: "divisors of 1 integer",
			args: args{x: []int{320}},
			want: nil,
		},
		{
			name: "common divisors of two integers",
			args: args{x: []int{320, 1500}},
			want: []int{1, 2, 4, 5, 10, 20},
		},
		{
			name: "common divisors of three integers",
			args: args{x: []int{320, 1500, 110}},
			want: []int{1, 2, 5, 10},
		},
		{
			name: "common divisors of big integers",
			args: args{x: []int{32000000, 15000000, 11000000, 12000000}},
			want: []int{1, 2, 4, 5, 8, 10, 16, 20, 25, 32, 40, 50,
				64, 80, 100, 125, 160, 200, 250, 320, 400, 500,
				625, 800, 1000, 1250, 1600, 2000, 2500, 3125,
				4000, 5000, 6250, 8000, 10000, 12500, 15625,
				20000, 25000, 31250, 40000, 50000, 62500, 100000,
				125000, 200000, 250000, 500000, 1000000},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			defer func() {
				end := time.Since(start)
				fmt.Println(end)
			}()

			if got := CommonDivisors(tt.args.x...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommonDivisors() = %v, want %v", got, tt.want)
			}
		})
	}
}
