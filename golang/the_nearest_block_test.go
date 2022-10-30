package main

import "testing"

func Test_theNearestBlock(t *testing.T) {
	type args struct {
		blocks []map[string]bool
		reqs   []string
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "test 1",
			args: args{
				blocks: []map[string]bool{
					{"gym": false, "school": true, "store": false},
					{"gym": true, "school": false, "store": false},
					{"gym": true, "school": true, "store": false},
					{"gym": false, "school": true, "store": false},
					{"gym": false, "school": true, "store": true},
				},
				reqs: []string{"gym", "school", "store"},
			},
			want: 3,
		},
		{
			name: "test 2",
			args: args{
				blocks: []map[string]bool{
					{"gym": false, "school": true, "store": false},
					{"gym": true, "school": false, "store": false},
					{"gym": true, "school": true, "store": false},
					{"gym": false, "school": true, "store": false},
					{"gym": false, "school": true, "store": true},
				},
				reqs: []string{"school", "store"},
			},
			want: 4,
		},
		{
			name: "test 3",
			args: args{
				blocks: []map[string]bool{
					{"gym": false, "school": true, "store": false},
					{"gym": true, "school": false, "store": false},
					{"gym": true, "school": true, "store": false},
					{"gym": false, "school": true, "store": false},
					{"gym": false, "school": true, "store": true},
				},
				reqs: []string{"gym"},
			},
			want: 1,
		},
		{
			name: "test 4",
			args: args{
				blocks: []map[string]bool{
					{"gym": false, "school": true, "store": false},
					{"gym": true, "school": false, "store": false},
					{"gym": true, "school": true, "store": false},
					{"gym": false, "school": true, "store": false},
					{"gym": false, "school": true, "store": true},
				},
				reqs: []string{"gym", "school"},
			},
			want: 2,
		},
		{
			name: "test 5",
			args: args{
				blocks: []map[string]bool{},
				reqs: []string{"gym", "school"},
			},
			want: 0,
		},
		{
			name: "test 6",
			args: args{
				blocks: []map[string]bool{
					{"gym": false, "school": true, "store": false},
					{"gym": true, "school": false, "store": false},
					{"gym": true, "school": true, "store": false},
					{"gym": false, "school": true, "store": false},
					{"gym": false, "school": true, "store": true},
				},
				reqs: []string{},
			},
			want: 0,
		},
		{
			name: "test 7",
			args: args{
				blocks: nil,
				reqs: []string{"gym", "school"},
			},
			want: 0,
		},
		{
			name: "test 8",
			args: args{
				blocks: []map[string]bool{
					{"gym": false, "school": true, "store": false},
					{"gym": true, "school": false, "store": false},
					{"gym": true, "school": true, "store": false},
					{"gym": false, "school": true, "store": false},
					{"gym": false, "school": true, "store": true},
				},
				reqs: nil,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := theNearestBlock(tt.args.blocks, tt.args.reqs); got != tt.want {
				t.Errorf("theNearestBlock() = %v, want %v", got, tt.want)
			}
		})
	}
}
