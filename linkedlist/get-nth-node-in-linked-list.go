package linkedlist

import "log"

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
