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

func shiftNodePosition(head *SinglyLinkedListNode, shift int32) *SinglyLinkedListNode {

	// if there's no node or no shifting than just return the original head
	if shift == 0 || head == nil {
		return head
	}

	// mapping the nodes
	nodeMap := map[int32]*SinglyLinkedListNode{}
	var count int32 = 1

	for head.next != nil {
		nodeMap[count] = head
		head = head.next
		count++
	}
	nodeMap[count] = head

	// find the modulo between shift number and the length of the node
	shift = shift % count

	// if no shifting just return the original
	if shift == 0 {
		return head
	}

	// adjust the next node
	nodeMap[count].next = nodeMap[1]
	nodeMap[count-shift].next = nil

	// change the head with the (length of the node - shift) + 1 node
	head = nodeMap[(count-shift)+1]

	return head
}

func deleteNode(llist *SinglyLinkedListNode, position int32) *SinglyLinkedListNode {
	if llist == nil {
		return nil
	}

	if position < 0 {
		return llist
	}

	if position == 0 {
		return llist.next
	}

	head := llist
	var currPos int32 = 1
	for head != nil {
		if currPos == position {
			head.next = head.next.next
			break
		}

		head = head.next
		currPos++
	}

	return llist
}

func reverseNode(llist *SinglyLinkedListNode) *SinglyLinkedListNode {
	if llist == nil {
		return nil
	}

	var curNode = llist
	var prevNode *SinglyLinkedListNode

	for curNode != nil {
		tempNext := curNode.next
		curNode.next = prevNode
		prevNode = curNode
		llist = curNode
		curNode = tempNext
	}

	return llist
}
