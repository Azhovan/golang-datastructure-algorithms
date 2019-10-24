package linkedlist

import (
	"fmt"
)

// Given a ‘key’, delete the first occurrence of this key in linked list
// linked list is not doubly
// Solution :
// find the element at key-1 (element(key-1))
// do element(key-1).next = element(key).next
// delete element(key)
func DeleteNodeAt(key int) {
	setup()

	if key == 0 {
		DeleteHead(key)
	} else {
		DeleteNoneHeadNodeAt(key)
	}

	list.decreaseLength()
	fmt.Printf("after deletion list size is: %d\n", list.len)
	list.IterateList()

}

func DeleteHead(key int) {
	// if it is head
	// introduce new element as head
	// remove old head from list
	newHead := list.headNode.nextNode
	list.headNode.nextNode = nil // remove reference to hel gc to remove it
	list.headNode = newHead
}

func DeleteNoneHeadNodeAt(key int) {
	// find elem in key-1
	node := list.FindByIndex(key - 1)
	next := node.nextNode.nextNode
	// remove the reference of deleted node
	node.nextNode.nextNode = nil
	// set the reference to new element
	node.nextNode = next
}

// delete all nodes from the beginning of the list to end
// time complexity is O(1)
func DeleteAllNodes() {
	setup()

	for i := 0; i < list.Len(); i++ {
		DeleteNodeAt(i)
	}
	list.IterateList()
}
