package linkedlist

import (
	"fmt"
	"strings"
)

// Solution
// hold start
// move the end pointer to next
// deduct in each step 1 from k
func RecursiveChunkLinkedList(start, end *Node, k int) {
	if k == 1 {
		// print it
		for s := start; s!= end.nextNode; s = s.nextNode {
			k++
			fmt.Println(s.property)
		}
		fmt.Println(strings.Repeat("-", 5))

		start = end.nextNode
	}

	// edge case
	if end == nil {
		fmt.Println(start.property)
		return
	}

	RecursiveChunkLinkedList(start, end.nextNode, k-1)
}



// create chunks with size k from linked list l
// Solution :
// every chunk has start, and size (k)
// loop over the list count the nodes
func NonRecursiveChunkLinkedList(k int) {
	setup()

	var start *Node
	i := 1

	for node := list.headNode; node != nil; node = node.nextNode {
		if i == 1 {
			start = node
		}

		if i == k {
			for ii := 0; ii < k; ii++ {
				fmt.Println(start.property)
				start = start.nextNode
			}
			// finish printing
			fmt.Println(strings.Repeat("-", 5))
			i = 1       // reset
			start = nil // reset
		} else {
			i++
		}
	}

}
