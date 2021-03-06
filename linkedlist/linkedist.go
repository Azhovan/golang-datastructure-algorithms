package linkedlist

import (
	"fmt"
	"log"
)

// this is very basic linked-ist implementation
// it could be generalized or implemented based on specific needs
type Node struct {
	property int
	nextNode *Node
}

func NewNode(p int) *Node {
	return &Node{
		property: p,
		nextNode: nil,
	}
}

func (n *Node) NextNode(next *Node) {
	n.nextNode = next
}

func (n *Node) Property() int{
	return n.property
}

type LinkedList struct {
	headNode *Node
	len      int // length
}

func (l *LinkedList) SetHead(n *Node) {
	l.headNode = n
}

func (l *LinkedList) increaseLength() {
	l.len++
}

func (l *LinkedList) decreaseLength() {
	l.len--
}

func (l *LinkedList) Len() int {
	return l.len
}

func (l *LinkedList) AddToHead(property int) {
	var node = &Node{
		property: property,
		nextNode: nil,
	}

	l.headNode = node
	l.increaseLength()
}

// iterate over the list and print the nodes
func (l *LinkedList) IterateList() {
	for node := l.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.property)
	}
	fmt.Println("-----------")
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
	l.increaseLength()
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
			l.increaseLength()
		}
	}
}

// return the node with property p
func (l *LinkedList) Get(p int) *Node {
	if l.Has(p) == false {
		log.Fatalf("node with value %d not found", p)
	}

	for node := l.headNode; node != nil; node = node.nextNode {
		if node.property == p {
			return node
		}
	}

	return nil
}

// find the node with index i
func (l *LinkedList) FindByIndex(index int) *Node {
	if l.Len() == 0 {
		log.Fatalf("list is empty")
	}
	var i = 0
	for node := l.headNode; node != nil; node = node.nextNode {
		if i == index {
			return node
		} else {
			i++
		}
	}
	return nil
}

// find the mean of the list
// notice that we are calculating index, which is started from 0

func (l *LinkedList) Mean() int {
	len := l.Len()

	if len%2 == 0 {
		i := l.FindByIndex(len/2 - 1).property
		j := l.FindByIndex(len / 2).property
		return (i + j) / 2
	} else {
		return l.FindByIndex(len / 2).property
	}
}
