package matrix

import "fmt"

func surfaceArea(A [][]int32) int32 {
	var area int32 = 0
	for r, rows := range A {
		for c := range rows {
			area += calculateArea(A, r, c)
		}
	}
	return area
}

func calculateArea(A [][]int32, r, c int) int32 {
	a := A[r][c]
	neighbor := findAdajacent(A, r, c)
	fmt.Println(neighbor)
	area := int32(2)
	for _, nArea := range neighbor {
		if nArea == 0 {
			area += a
		} else if a > nArea {
			area += a - nArea
		}
	}
	return area
}

func findAdajacent(A [][]int32, r, c int) []int32 {
	var res = make([]int32, 4)

	valid := func(r, c int) bool {
		if r < 0 || r >= len(A) || c < 0 || c >= len(A[0]) {
			return false
		}
		return true
	}

	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for i, d := range directions {
		if valid(r+d[0], c+d[1]) {
			res[i] = A[r+d[0]][c+d[1]]
		}
	}

	return res
}
