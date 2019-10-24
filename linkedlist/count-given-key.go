package linkedlist


//Given a singly linked list and a key,
// count number of occurrences of all keys in linked list.
// For example,
// if given linked list is 1->2->1->2->1->3->1, then [1:4, 2:2, 3:1]
// then output should be 4.
func CountKeys() map[int]int {
	setup()
	counts := make(map[int]int)
	for node := list.headNode; node != nil; node = node.nextNode {
		if counts[node.property] != 0 {
			counts[node.property] += 1
		} else {
			counts[node.property] = 1
		}
	}
	return counts
}


//Given a singly linked list and a key,
// count number of occurrences of given key in linked list.
// For example,
// if given linked list is 1->2->1->2->1->3->1 and given key is 1,
// then output should be 4.

func CountKey(p int) int {
	setup()
	i := 0
	for node := list.headNode; node != nil; node = node.nextNode {
		if node.property == p {
			i++
		}
	}
	return i
}

