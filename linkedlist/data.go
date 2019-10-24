package linkedlist

import (
	"fmt"
)

var list LinkedList

func setup() {
	list.AddToHead(0)
	list.AddToEnd(1)
	list.AddToEnd(2)
	list.AddToEnd(3)
	list.AddToEnd(5)
	list.AddToEnd(4)
	list.AddToEnd(5)
	list.AddToEnd(5)
	list.AddToEnd(5)

	fmt.Println("created list: ")
	list.IterateList()
}

