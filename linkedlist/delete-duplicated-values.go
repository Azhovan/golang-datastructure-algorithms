package linkedlist

//You're given the pointer to the head node of a sorted linked list,
//where the data in the nodes is in ascending order.
//Delete as few nodes as possible so that the list does not contain any value more than once.
//The given head pointer may be null indicating that the list is empty.

// example
// The initial linked list is: 1 -> 2 -> 2 -> 3 -> 4 -> NULL
//The final linked list is: 1 -> 2 -> 3 -> 4 -> NULL
func RemoveDuplicates() {
	setup()
	var hashmap = make(map[int]bool)
	var prev *Node

	for node := list.headNode; node != nil; node = node.nextNode {
		_, exist := hashmap[node.property]
		if exist == false {
			hashmap[node.property] = true
			prev = node
		} else {
			prev.nextNode = node.nextNode
		}
	}

	list.IterateList()
}
