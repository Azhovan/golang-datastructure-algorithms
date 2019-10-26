package set

import "testing"

func TestSet_New(t *testing.T) {
	var s = &Set{}
	s.New()
	if s.Get() == nil {
		t.Error("invalid definition")
	}
}

func TestSet_AddElement(t *testing.T) {
	var s = &Set{}
	s.New()
	s.AddElement(1)
	s.AddElement(2)
	s.AddElement(3)

	if len(s.Get()) != 3 {
		t.Error("Set s should have 3 items")
	}
}

func TestSet_Contains(t *testing.T) {
	var s = &Set{}
	s.New()
	s.AddElement(1)
	s.AddElement(2)
	s.AddElement(3)

	if s.Contains(2) == false {
		t.Error("given Set s, should contains element with value of 2")
	}
}

func TestSet_DeleteElement(t *testing.T) {
	var s = &Set{}
	s.New()
	s.AddElement(1)
	s.AddElement(2)
	s.AddElement(3)

	s.DeleteElement(2)

	if s.Contains(2) == true {
		t.Error("given Set s, should not contains element with value of  2")
	}
}

func TestSet_Get(t *testing.T) {
	var s = &Set{}
	s.New()
	s.AddElement(1)
	s.AddElement(2)
	s.AddElement(3)

	collection := s.Get()
	if collection[1] == false || collection[2] == false || collection[3] == false {
		t.Error("invalid elements in Set s")
	}
}

func TestSet_Union(t *testing.T) {
	var s1 = &Set{}
	s1.New()
	s1.AddElement(1)
	s1.AddElement(2)
	s1.AddElement(3)
	s1.AddElement(4)

	var s2 = &Set{}
	s2.New()
	s2.AddElement(5)
	s2.AddElement(6)
	s2.AddElement(7)

	union := s1.Union(s2)
	if union.Contains(1) == false ||
		union.Contains(2) == false ||
		union.Contains(3) == false ||
		union.Contains(4) == false ||
		union.Contains(5) == false ||
		union.Contains(6) == false ||
		union.Contains(7) == false {
		t.Error("wrong union operation")
	}

}

func TestSet_Intersect(t *testing.T) {
	var s1 = &Set{}
	s1.New()
	s1.AddElement(1)
	s1.AddElement(2)
	s1.AddElement(3)
	s1.AddElement(4)

	var s2 = &Set{}
	s2.New()
	s2.AddElement(3)
	s2.AddElement(2)
	s2.AddElement(7)
	s2.AddElement(8)

	intersect := s2.Intersect(s1)

	if intersect.Contains(2) == false || intersect.Contains(3) == false {
		t.Error("wrong intersect operation")
	}

	if intersect.Contains(1) == true ||
		intersect.Contains(4) == true ||
		intersect.Contains(7) == true ||
		intersect.Contains(8) == true {
		t.Error("wrong union operation")
	}
}
