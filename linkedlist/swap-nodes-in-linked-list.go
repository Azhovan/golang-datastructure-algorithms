package linkedlist

// Difficulty: medium

//Given a linked list and two keys in it, swap nodes for
//two given keys. Nodes should be swapped by changing links.
//	Swapping data of nodes may be expensive in many situations when data contains many fields.
//It may be assumed that all keys in linked list are distinct.

//Input:  10->15->12->13->20->14,  x = 12, y = 20
//Output: 10->15->20->13->12->14

func SwapNodes(i, j int) {
	setup()

	// find the element(p-1)
	p0 := list.FindByIndex(i-1)
	p1 := p0.nextNode

	n0 := list.FindByIndex(j-1)
	n1 := n0.nextNode

	// From ....p0->p1.....n0->n1...
	// To ....p0->n1->.....n0->p1->...

	p0.nextNode = n1
	n0.nextNode = p1

	oldP1next := p1.nextNode
	oldN1next := n1.nextNode

	n1.nextNode = oldP1next
	p1.nextNode = oldN1next

	list.IterateList()
}