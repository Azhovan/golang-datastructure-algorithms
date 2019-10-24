package linkedlist

// Given a linked list, check if the linked list has loop or not.
// Below diagram shows a linked list with a loop.

func DetectLoop(list LinkedList) bool {

	hashshmap := make(map[*Node]bool)
	for node := list.headNode; node != nil; node = node.nextNode {
		if hashshmap[node] == true {
			return true
		} else {
			hashshmap[node] = true
		}
	}

	return false
}
