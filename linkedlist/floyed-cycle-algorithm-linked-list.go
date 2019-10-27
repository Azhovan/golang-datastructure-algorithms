package linkedlist

// Floyd algorithm
// detect cycle in linked list
func Floyd_cycle_detection_algorithm(list LinkedList) (bool, *Node) {
	trutle := list.headNode
	hare := list.headNode
	var found bool
	for {
		if hare == nil {
			break
		}

		hare = hare.nextNode.nextNode
		trutle = trutle.nextNode

		if hare == trutle {
			found = true
			break
		}
	}
	return found, trutle
}

//  Detect The First Node of The Cycle
func Floyd_first_node_of_cycle(list LinkedList) *Node {
	found, meetingNode := Floyd_cycle_detection_algorithm(list)
	if found == false {
		return nil
	}

	t1 := list.headNode
	t2 := meetingNode

	var firstNodeOfCycle *Node
	for {
		if t1 == t2 {
			firstNodeOfCycle = t1
			break
		}

		t1 = t1.nextNode
		t2 = t2.nextNode
	}

	return firstNodeOfCycle
}
