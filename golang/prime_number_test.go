package main

import (
	"reflect"
	"testing"
)

func TestPrimeNumber(t *testing.T) {
	type args struct {
		limit int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "prime number 1 to 100",
			args: args{limit: 100},
			want: []int{2, 3, 4, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrimeNumber(tt.args.limit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PrimeNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
