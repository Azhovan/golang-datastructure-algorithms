package linkedlist


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
