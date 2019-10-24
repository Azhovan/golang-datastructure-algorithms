package linkedlist

import (
	"fmt"
	"log"
)

//Given pointer to the head node of a linked list,
// the task is to reverse the linked list.
// We need to reverse the list by changing links between nodes.
// none recursive way
func Revers() {
	setup()
	if list.Len() == 0 {
		log.Fatalf("list is empty")
	}
	if list.Len() == 1 {
		fmt.Print(list.headNode.property)
		return
	}

	current := list.headNode
	next := current.nextNode

	list.headNode = list.LastNode()

	reversList(nil, current, next)
	list.IterateList()
}

func reversList(prev, current, next *Node) {

	if current == nil {
		return
	}

	current.nextNode = prev

	if next == nil {
		return
	}
	reversList(current, next, next.nextNode)
}
