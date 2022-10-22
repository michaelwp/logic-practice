package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
	"testing"
)

func Test_printLinkedList(t *testing.T) {
	type args struct {
		head *SinglyLinkedListNode
	}
	tests := []struct {
		name   string
		args   args
		expect []string
	}{
		{
			name: "linked_list",
			args: args{head: &SinglyLinkedListNode{
				data: 1,
				next: &SinglyLinkedListNode{
					data: 2,
					next: &SinglyLinkedListNode{
						data: 3,
						next: &SinglyLinkedListNode{
							data: 4,
							next: nil,
						},
					},
				},
			}},
			expect: []string{"1", "2", "3", "4", ""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			expect := tt.expect
			actual := strings.Split(captureStdout(tt.args.head), "\n")

			if !reflect.DeepEqual(actual, expect) {
				fmt.Printf("expect: %s\nactual: %s", expect, actual)
			}
		})
	}
}

func captureStdout(head *SinglyLinkedListNode) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	printLinkedList(head)

	w.Close()
	os.Stdout = old

	outChan := make(chan string)
	defer close(outChan)
	go func(outChan chan string) {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outChan <- buf.String()
	}(outChan)

	return <-outChan
}
