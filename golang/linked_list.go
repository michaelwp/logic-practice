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

func insertNodeAtPosition(llist *SinglyLinkedListNode, data int32, position int32) *SinglyLinkedListNode {
	var i int32 = 0
	var tempLlist = llist
	var tempNext *SinglyLinkedListNode

	for tempLlist != nil {
		if i == position {
			tempLlist = &SinglyLinkedListNode{
				data: data,
				next: tempLlist,
			}
		}

		if tempNext != nil {
			tempNext.next = tempLlist
		}
		tempNext = tempLlist
		tempLlist = tempLlist.next

		i++
	}

	return llist
}
