package main

import (
	"reflect"
	"testing"
)

func Test_numberOfOccurrences(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name    string
		args    args
		wantRes map[string]int
	}{
		{
			name: "finaccel",
			args: args{word: "Finaccel"},
			wantRes: map[string]int{"a":1, "c":2, "e":1, "f":1, "i":1, "l":1, "n":1},
		},
		{
			name: "test",
			args: args{word: "Test"},
			wantRes: map[string]int{"e":1, "s":1, "t":2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := numberOfOccurrences(tt.args.word); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("numberOfOccurrences() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
