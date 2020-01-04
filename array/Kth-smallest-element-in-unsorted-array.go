package array

import "container/heap"

// K’th Smallest Element in Unsorted Array
// Given an array and a number k where k is smaller than size of array,
// we need to find the k’th smallest element in the given array.

type IntHeap []int

// How to use:
// 	list := array.IntHeap{2, 6, 7, 3, 9, 5, 99}
//	kth := array.KthSmallest(list, 6)
//	fmt.Println(kth) // will print 9
func KthSmallest(A IntHeap, k int) interface{} {
	if k > A.Len() {
		return nil
	}

	heap.Init(&A)
	var kth interface{}
	for i := 0; i < k; i++ {
		kth = heap.Pop(&A)
	}
	return kth
}

// returns heap size
func (h IntHeap) Len() int {
	return len(h)
}

// compare two elements in heap
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

// swap two indexes in heap
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// push the new elements into heap
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// pop the smallest element from the heap
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
