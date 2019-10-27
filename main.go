package main

import (
	"fmt"
	"hackerrank-solutions/linkedlist"
)

func main() {

}


// example
func detect_cycle_floyd_algorithm() {
	n4 := linkedlist.NewNode(5)
	n3 := linkedlist.NewNode(4)
	n2 := linkedlist.NewNode(3)
	n1 := linkedlist.NewNode(2)
	n0 := linkedlist.NewNode(1)
	n0.NextNode(n1)
	n1.NextNode(n0) // loop here !
	n2.NextNode(n3)
	n3.NextNode(n4)
	n4.NextNode(n1)

	var l linkedlist.LinkedList
	l.SetHead(n0)

	found, meetingNode := linkedlist.Floyd_cycle_detection_algorithm(l)
	fmt.Print("cycle : ", found)
	if found == true {
		fmt.Print("the first node of the cycle has ", meetingNode.Property(), " value")
	}

}

// example
func find_loop_example() {
	n4 := linkedlist.NewNode(5)
	n3 := linkedlist.NewNode(4)
	n2 := linkedlist.NewNode(3)
	n1 := linkedlist.NewNode(2)
	n0 := linkedlist.NewNode(1)
	n0.NextNode(n1)
	n1.NextNode(n0) // loop here !
	n2.NextNode(n3)
	n3.NextNode(n4)
	n4.NextNode(n1)

	var l linkedlist.LinkedList
	l.SetHead(n0)

	fmt.Print(linkedlist.DetectLoop(l))
}

// example
func deleteNodes_example() {
	linkedlist.DeleteNodeAt(0)
	linkedlist.DeleteNodeAt(1)
	linkedlist.DeleteAllNodes()
}

// example
func other_examples() {
	// fmt.Print(linkedlist.Length()) // 5
	// fmt.Print(linkedlist.NthFromEnd(4))
	// linkedlist.Revers()
	// linkedlist.SwapNodes(1,3)
	// fmt.Println(linkedlist.CountKey(5))
	//linkedlist.DeleteNodeAt(0)
	// linkedlist.RemoveDuplicates()
}

func count_all_keys_occurrence_examples() {
	for i, v := range linkedlist.CountKeys() {
		fmt.Printf("for key %d, occurrence is: %d\n", i, v)
	}
}
