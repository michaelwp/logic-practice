package main

import "fmt"

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func printLinkedList(head *SinglyLinkedListNode) {
	for head.next != nil {
		fmt.Println(head.data)
		head = head.next
	}

	fmt.Println(head.data)
}
