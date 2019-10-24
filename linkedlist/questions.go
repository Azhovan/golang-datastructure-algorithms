package linkedlist

import (
	"fmt"
	"log"
)

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

// recursive length calculations
func Length() int {
	setup()
	return lengthRecursive(list.headNode)
}

func lengthRecursive(node *Node) int {
	if node.nextNode == nil {
		return 1
	}
	return lengthRecursive(node.nextNode) + 1
}

// Given a Linked List and a number n,
// write a function that returns the value at
// the n’th node from the end of the Linked List.
func NthFromEnd(i int) int {
	setup()
	if i > list.Len() || i < 0 {
		log.Fatalf("index is not valid, given: %d but the list length is :%d", i, list.Len())
	}

	index := list.Len() - i
	return list.FindByIndex(index).property
}

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
