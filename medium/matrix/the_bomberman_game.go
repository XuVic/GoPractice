package matrix

var grid = []string{".......", "...O...", "....O..", ".......", "OO.....", "OO....."}
var second int = 3

var grid2 = []string{
	".......",
	"...O.O.",
	"....O..",
	"..O....",
	"OO...OO",
	"OO.O...",
}

func bomberMan(n int32, grid []string) []string {
	fullGrid := make([]string, len(grid))
	for i := range fullGrid {
		for j := 0; j < len(grid[0]); j++ {
			fullGrid[i] += "O"
		}
	}

	if n == 1 {
		return grid
	} else if n == 2 {
		return fullGrid
	} else if n == 3 {
		return reverse(grid)
	}

	grid = reverse(grid)
	n = n - 3
	action := n % 2

	if action == 1 {
		return fullGrid
	}

	bombTime := int(n / 2)
	if bombTime%2 == 1 {
		return reverse(grid)
	}
	return grid
}

func reverse(grid []string) []string {

	res := make([]string, len(grid))

	for r, rows := range grid {
		for c := range rows {
			if valid(grid, r, c) {
				res[r] += "O"
			} else {
				res[r] += "."
			}
		}
	}
	return res

}

func valid(grid []string, r int, c int) bool {
	if grid[r][c] == 'O' {
		return false
	}

	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for _, d := range directions {
		newR := r + d[0]
		newC := c + d[1]

		if newR >= 0 && newR < len(grid) && newC >= 0 && newC < len(grid[0]) {
			if grid[newR][newC] == 'O' {
				return false
			}
		}
	}

	return true
}
