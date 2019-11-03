package linkedlist

import (
	"fmt"
	"strings"
)

//  chunk linked list none recursive
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
