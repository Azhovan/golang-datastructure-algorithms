package heap

/// heap elements
type heap []int

func NewMaxHeap() *heap {
	return &heap{}
}

// add new element to heap
func (h *heap) Push(value int) {
	*h = append(*h, value)
	h.up(h.Len() - 1)
}

func (h *heap) up(i int) {
	for {
		j := (i - 1) / 2
		if i == j || (*h)[j] > (*h)[i] {
			break
		}
		h.Swap(i, j)
		i = j
	}
}

// remove the root
// replace it with last element
func (h *heap) Pop() int {
	top := (*h)[0]
	// instead of shifting the array,
	//we just swap the last and first element
	h.Swap(0, h.Len()-1)
	*h = (*h)[:h.Len()-1]
	h.down(0)
	return top
}

func (h *heap) down(i int) {
	for {
		// there is more chance to have left child
		// than right child, so first we compute it
		l := 2*i + 1
		if l >= h.Len() || l < 0 {
			break
		}
		j := l
		// now check the right child
		if r := l + 1; r < h.Len() && (*h)[r] > (*h)[l] {
			j = r
		}
		if (*h)[i] > (*h)[j] {
			break
		}
		h.Swap(i, j)
		i = j
	}
}

func (h *heap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

// size of the elements in heap
func (h *heap) Len() int {
	return len(*h)
}

// Lists the heap elements
func (h *heap) List() heap {
	return *h
}
