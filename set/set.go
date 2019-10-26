package set

import "log"

type Set struct {
	collection map[int]bool
}

func (s *Set) New() {
	s.collection = make(map[int]bool)
}

func (s *Set) Get() map[int]bool {
	return s.collection
}

// add element p to Set s if the element p does not exist
func (s *Set) AddElement(p int) {
	if s.Contains(p) == false {
		s.collection[p] = true
	}
}

// check if an item with type integer exists in Set s
func (s *Set) Contains(p int) bool {
	var exists bool
	_, exists = s.collection[p]
	return exists
}

// delete element p from Set s if it does exists
func (s *Set) DeleteElement(p int) {
	if s.Contains(p) == false {
		log.Fatalf("given element %d, does not exists", p)
	}
	delete(s.collection, p)
}

// return intersect of given Sets of s1 and s2
func (s1 *Set) Intersect(s2 *Set) *Set {
	var intersect = &Set{}
	intersect.New()

	// index holds items
	for i, _ := range s1.collection {
		if s2.Contains(i) {
			intersect.AddElement(i)
		}
	}

	return intersect
}

// return union of given Sets s1 and s2
func (s1 *Set) Union(s2 *Set) *Set {
	var unionset = &Set{}
	unionset.New()

	for i, _ := range s1.collection {
		unionset.AddElement(i)
	}

	for j, _ := range s2.collection {
		unionset.AddElement(j)
	}

	return unionset
}
