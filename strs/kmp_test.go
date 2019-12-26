package strs

import "testing"

func TestLspTable(t *testing.T) {
	expected := []int{0, 0, 0, 1, 2, 1, 0, 0, 0}
	pattern := "acbacabcy"
	for i , v := range LspTable(pattern) {
		if v != expected[i] {
			t.Errorf("wrong lsp table")
		}
	}
}

func TestFindExistedPattern(t *testing.T) {
	pattern := "acbacabcy"
	text := "acbacabcyzacbacabcyxacbacabcy"

	if FindPattern(text, pattern) != true {
		t.Errorf("patern not found")
	}
}

func TestNotExistedPattern(t *testing.T) {
	pattern := "acbacabcb"
	text := "acbacabcyzacbacabcyxacbacabcy"

	if FindPattern(text, pattern) == true {
		t.Errorf("wrong pattern match")
	}
}
