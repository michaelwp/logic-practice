package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_minDifference(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{
				a: []int{0, 1, 1, 0, 0, 1},
				b: []int{10, 2, 3, 1, 7, 18},
			},
			want: 1,
		},
		{
			name: "test 2",
			args: args{
				a: []int{0, 1, 1, 0, 0, 1, 0, 1, 1, 0, 0, 0},
				b: []int{10, 2, 3, 1, 17, 18, 10, 22, 3, 31, 27, 18},
			},
			want: 0,
		},
		{
			name: "test 2",
			args: args{
				a: []int{0, 1, 1, 0, 0, 1, 0, 1, 1, 0, 0, 0, 1, 1, 0},
				b: []int{10, 200, 3, 1, 17, 108, 1000, 22, 3, 310, 127, 18, 998, 1000000, 999999},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			defer func() {
				end := time.Since(start)
				fmt.Println(end)
			}()

			if got := MinDifference(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("minDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkMinDifference(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := []int{0, 1, 1, 0, 0, 1, 0, 1, 1, 0, 0, 0, 0, 1, 1, 0, 0, 1, 0, 1, 1, 0, 0, 0,
			0, 1, 1, 0, 0, 1, 0, 1, 1, 0, 0, 0, 0, 1, 1, 0, 0, 1, 0, 1, 1, 0, 0, 0,
			0, 1, 1, 0, 0, 1, 0, 1, 1, 0, 0, 0, 0, 1, 1, 0, 0, 1, 0, 1, 1, 0, 0, 0,
			0, 1, 1, 0, 0, 1, 0, 1, 1, 0, 0, 0, 0, 1, 1, 0, 0, 1, 0, 1, 1, 0, 0, 0,
			0, 1, 1, 0, 0, 1, 0, 1, 1, 0, 0, 0, 0, 1, 1, 0, 0, 1, 0, 1, 1, 0, 0, 0}
		b := []int{10, 2, 3, 1, 17, 18, 10, 22, 3, 31, 27, 18, 10, 2, 3, 1, 17, 18, 10, 22, 3, 31, 27, 18,
			10, 2, 3, 1, 17, 18, 10, 22, 3, 31, 27, 18, 10, 2, 3, 1, 17, 18, 10, 22, 3, 31, 27, 18,
			10, 2, 3, 1, 17, 18, 10, 22, 3, 31, 27, 18, 10, 2, 3, 1, 17, 18, 10, 22, 3, 31, 27, 18,
			10, 2, 3, 1, 17, 18, 10, 22, 3, 31, 27, 18, 10, 2, 3, 1, 17, 18, 10, 22, 3, 31, 27, 18,
			10, 2, 3, 1, 17, 18, 10, 22, 3, 31, 27, 18, 10, 2, 3, 1, 17, 18, 10, 22, 3, 31, 27, 18}

		MinDifference(a, b)
	}
}
