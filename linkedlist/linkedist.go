package linkedlist

import "fmt"

// this is very basic linked-ist implementation
// it could be generalized or implemented based on specific needs
type Node struct {
	property int
	nextNode *Node
}

type LinkedList struct {
	headNode *Node
}

func (l *LinkedList) AddToHead(property int) {
	var node = &Node{
		property: property,
		nextNode: nil,
	}

	l.headNode = node
}

// iterate over the list and print the nodes
func (l *LinkedList) IterateList() {
	for node := l.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.property)
	}
}

// print the the last node value
func (l *LinkedList) LastNode() *Node {
	var node *Node

	// if there is no predefined head, return nil
	if l.headNode == nil {
		return nil
	}

	for node = l.headNode; node.nextNode != nil; {
		node = node.nextNode
	}
	return node
}

// Add the node to the end of the list
func (l *LinkedList) AddToEnd(property int) {
	lastNode := l.LastNode()
	var newNode = &Node{
		property: property,
		nextNode: nil,
	}
	lastNode.nextNode = newNode
}

// return true if there will be at least one occurrence of the p in LinkedList l
func (l *LinkedList) Has(p int) bool {
	for node := l.headNode; node != nil; node = node.nextNode {
		if node.property == p {
			return true
		}
	}
	return false
}

// return the index of the first occurrence of the p in LinkedList l
// index is zero based
// -1 means no Node with p exist
func (l *LinkedList) Index(p int) int {
	var i int
	for node := l.headNode; node != nil; node = node.nextNode {
		if node.property == p {
			return i
		}
		i++
	}
	return -1
}

// add the Node n after Node with property p
// if p did not exist, error returned
// new position of the n will be returned
func (l *LinkedList) AddAfter(p int, n int) {
	if l.Has(p) == false {
		fmt.Errorf("node with property %d doesn't exist", p)
	}
	for node := l.headNode; node != nil; node = node.nextNode {
		if node.property == p {
			newNode := &Node{
				property: n,
				nextNode: node.nextNode,
			}
			node.nextNode = newNode
		}
	}
}

// return the node with property p
func (l *LinkedList) Get(p int) *Node {
	if l.Has(p) == false {
		fmt.Errorf("node with value %d not found", p)
	}

	for node := l.headNode; node != nil; node = node.nextNode {
		if node.property == p {
			return node
		}
	}

	return nil
}
