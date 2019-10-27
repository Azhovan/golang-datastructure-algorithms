package queues

// every queue can store specific type to store
// in this case we have slices of Items
type Queue []Items

// this is simple object which contains only priority and value
// priority will indicate how much this item is important to enqueue
// typically high priorities items will be proceed earlier before others
type Items struct {
	priority int
	value    int
}

// create a new Item from inputs
func (i *Items) New(priority, value int) {
	i.priority = priority
	i.value = value
}

func (i *Items) Data() int {
	return i.value
}

func (i *Items) Priority() int {
	return i.priority
}

// return the length of queue
func (q *Queue) Len() int {
	return len(*q)
}

// Add item type of Item to end of the queue with type of Queue
// new item will be added before the items with highest priority in the list
func (q *Queue) AddItem(item *Items) {
	appended := false
	for i, v := range *q {
		if item.priority >= v.priority {
			*q = append((*q)[:i], append(Queue{*item}, (*q)[i:]...)...)
			appended = true
			break
		}
	}

	if appended == false {
		*q = append(*q, *item)
	}
}
