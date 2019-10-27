package queues

import (
	"testing"
)

func TestItems_Priority(t *testing.T) {
	var item Items
	item.New(1, 2)

	if item.priority != 1 {
		t.Errorf("invalid priority, given %d", item.priority)
	}
}

func TestItems_Data(t *testing.T) {
	var item Items
	item.New(1, 2)

	if item.Data() != 2 {
		t.Errorf("invalid data, given %d", item.value)
	}
}

func TestQueue_AddItem(t *testing.T) {
	var item1 = &Items{}
	item1.New(1, 11)

	var item2 = &Items{}
	item2.New(2, 22)

	var q Queue
	q = make(Queue, 0)

	q.AddItem(item1)
	q.AddItem(item2)

	var item3 = &Items{}
	item3.New(3,33)

	q.AddItem(item3)

	if q.Len() != 3 {
		t.Errorf("length is wrong, given %d", q.Len())
	}

	if q[0].priority != 3 {
		t.Errorf("the first priority in the queue should be 3, given %d", q[0].priority)
	}

	if q[1].priority != 2 {
		t.Errorf("the first priority in the queue should be 2, given %d", q[0].priority)
	}

	if q[2].priority != 1 {
		t.Errorf("the first priority in the queue should be 1, given %d", q[0].priority)
	}
}
