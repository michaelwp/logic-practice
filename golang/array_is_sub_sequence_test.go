package main

import "testing"

func Test_isSubSequence(t *testing.T) {
	type args struct {
		arr  []int32
		sArr []int32
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test 1",
			args: args{
				arr:  []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
				sArr: []int32{3, 6, 0},
			},
			want: true,
		},{
			name: "test 2",
			args: args{
				arr:  []int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				sArr: []int32{3, 6, 0},
			},
			want: false,
		},{
			name: "test 3",
			args: args{
				arr:  []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
				sArr: []int32{},
			},
			want: false,
		},{
			name: "test 4",
			args: args{
				arr:  []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
				sArr: nil,
			},
			want: false,
		},{
			name: "test 5",
			args: args{
				arr:  nil,
				sArr: []int32{3, 6, 0},
			},
			want: false,
		},{
			name: "test 6",
			args: args{
				arr:  []int32{},
				sArr: []int32{3, 6, 0},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubSequence(tt.args.arr, tt.args.sArr); got != tt.want {
				t.Errorf("isSubSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
