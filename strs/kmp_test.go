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

func TestFindExistedPattern2(t *testing.T) {
	pattern := "ghi"
	text := "abcdefghabcdefghi"

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

func TestFindIndex(t *testing.T) {
	pattern := "st"
	text := "sst"

	if FindIndex(text, pattern) != 1 {
		t.Errorf("wrong index")
	}

}
