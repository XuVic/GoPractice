package matrix

import "testing"

var a = [][]int32{
	{1, 3, 4}, {2, 2, 3}, {1, 2, 4},
}

func TestAlog(t *testing.T) {
	t.Log(reverse(reverse(reverse(grid2))))
	t.Log(bomberMan(3, grid))
}
