package search

import (
	"strconv"
)

func gridSearch(G []string, P []string) string {

	p := P[0]

	for r, g := range G {
		if len(G)-r < len(P) {
			break
		}

		var c int
		var match bool
		c, match = horspoolSearch(g, p, len(P[0])-1)
		for match {
			if checkRest(G, P, r, c) {
				return "YES"
			}
			c, match = horspoolSearch(g, p, c+1)
		}
	}

	return "NO"
}

func checkRest(G []string, P []string, r, c int) bool {
	pLen := len(P[0])
	for i := 1; i < len(P); i++ {
		if G[r+i][c-pLen+1:c+1] != P[i] {
			return false
		}
	}
	return true
}

func horspoolSearch(text string, pattern string, index int) (int, bool) {
	table := createTable(pattern)

	m := len(pattern) - 1
	found := false
	posR := index
	for posR < len(text) {
		k := 0
		for k <= m && text[posR-k] == pattern[m-k] {
			k++
		}
		if k == m+1 {
			found = true
			break
		}
		posR += table[string(text[posR])]
	}
	return posR, found
}

func createTable(pattern string) map[string]int {
	table := make(map[string]int)
	for i := 0; i <= 9; i++ {
		table[strconv.Itoa(i)] = len(pattern)
	}
	for i := 0; i < len(pattern)-1; i++ {
		table[string(pattern[i])] = len(pattern) - 1 - i
	}
	return table
}
