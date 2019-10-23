package linkedlist

import "fmt"

var list LinkedList

func setup() {
	list.AddToHead(0)
	list.AddToEnd(1)
	list.AddToEnd(2)
	list.AddToEnd(3)
	list.AddToEnd(4)

	fmt.Println("list created.")
	list.IterateList()
}

// Given a ‘key’, delete the first occurrence of this key in linked list
// linked list is not doubly
// Solution :
// find the element at key-1 (element(key-1))
// do element(key-1).next = element(key).next
// delete element(key)
func DeleteNodeAt(key int) {
	setup()

	// find elem in key-1
	node := list.FindByIndex(key-1)
	next := node.nextNode.nextNode
	// remove the reference of deleted node
	node.nextNode.nextNode = nil
	// set the reference to new element
	node.nextNode = next
	list.decreaseLength()

	fmt.Printf("after deletion list size is: %d\n", list.len)
	list.IterateList()
}
