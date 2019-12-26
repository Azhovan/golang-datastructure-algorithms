package strs

// Knuth-Morris-Pratt KMP String Matching Algorithm

// prepare the the lsp table
func LspTable(pat string) []int {
	i := 1
	j := 0

	lsp := make([]int, len(pat))
	lsp[0] = 0

	for ; i < len(pat); {
		if pat[i] == pat[j] {
			lsp[i] = j + 1
			i++
			j++
			continue
		} else {
			j--
			if j < 0 {
				j = 0
				lsp[i] = 0
				i++
				continue
			}
			j = lsp[j]
			continue
		}
	}

	return lsp
}

func FindPattern(text, pat string) bool {
	lsp := LspTable(pat)
	i, j := 0, 0
	found := false

	for ; i < len(text) && j <len(pat); {
		if text[i] == pat[j] {
			if j == len(pat)-1 {
				found = true
				break
			}
			i++
			j++
		} else {
			j--
			if j < 0 {
				j = 0
				i++
				continue
			}
			j = lsp[j]
		}
	}
	return found
}

