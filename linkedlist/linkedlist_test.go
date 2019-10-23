package linkedlist

import (
	"testing"
)

func TestLinkedList_Has(t *testing.T) {
	var list LinkedList
	list.AddToHead(10)

	if b := list.Has(10); b == false {
		t.Errorf("list should has value 10")
	}
}

func TestLinkedList_Index(t *testing.T) {
	var list LinkedList

	list.AddToHead(10)
	list.AddAfter(10, 11)

	if list.Index(11) != 1 {
		t.Errorf("list should has value 11 at index 1")
	}
}

func TestLinkedList_AddAfter(t *testing.T) {
	var list LinkedList

	list.AddToHead(10)
	list.AddToEnd(11)
	list.AddToEnd(12)
	list.AddToEnd(13)

	list.AddAfter(13, 15)

	n := list.Get(13)

	if n == nil {
		t.Errorf("node with value 13 not found")
	}

	if n.nextNode != list.Get(15) {
		t.Errorf("next node after 13 is 15")
	}

}

func TestLinkedList_AddToEnd(t *testing.T) {
	var list LinkedList

	list.AddToHead(10)
	list.AddToEnd(11)

	if b := list.LastNode(); b.property != 11 {
		t.Errorf("list should has value 11")
	}
}

func TestLinkedList_Get(t *testing.T) {
	var list LinkedList

	list.AddToHead(10)
	list.AddToEnd(11)

	n := list.Get(11)
	if n.property != 11 {
		t.Errorf("value of the node should be 11")
	}
}

func TestLinkedList_LastNode(t *testing.T) {
	var list LinkedList

	list.AddToHead(10)
	list.AddToEnd(11)

	if b := list.LastNode(); b.property != 11 {
		t.Errorf("list should has value 11")
	}
}

func TestLinkedList_AddToHead(t *testing.T) {
	var list LinkedList
	list.AddToHead(10)

	if b := list.LastNode(); b.property != 10 {
		t.Errorf("list should has value 11")
	}

	var list2 LinkedList
	if b := list2.LastNode(); b != nil {
		t.Errorf("list2 does not have head yet")
	}

}

func TestLinkedList_Len(t *testing.T) {
	var l LinkedList
	l.AddToHead(2)
	l.AddToEnd(3)

	if l.Len() != 2 {
		t.Errorf("length of the linkedlist should be 2, given %d", l.len)
	}
}

func TestLinkedList_FindByIndex(t *testing.T) {
	var l LinkedList
	l.AddToHead(2)
	l.AddToEnd(3)

	if n := l.FindByIndex(1).property; n != 3 {
		t.Errorf("element at index 1 is 3, given %d", n)
	}
}

func TestLinkedList_Mid(t *testing.T) {
	var l LinkedList

	l.AddToHead(2)
	l.AddToEnd(3)
	l.AddToEnd(5)

	if n := l.Mean(); n != 3 {
		t.Errorf("mid value is 3. given %d", n)
	}

	l.AddToEnd(6)

	if n := l.Mean(); n != 4 {
		t.Errorf("mid value is 4. given %d", n)
	}
}
