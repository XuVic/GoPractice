package graph

func maxSize(m [][]int32) int32 {
	var max int32 = 0

	for r := int32(0); r < int32(len(m)); r++ {
		for c := int32(0); c < int32(len(m[0])); c++ {
			if m[r][c] == 1 {
				size := new(int32)
				*(size)++
				bfs(m, []int32{r, c}, size)
				if *(size) > max {
					max = *(size)
				}
			}
		}
	}
	return max
}

func bfs(m [][]int32, node []int32, size *int32) {
	r := node[0]
	c := node[1]
	m[r][c] = 2
	nodes := getAdjacet(m, node)
	for _, n := range nodes {
		m[n[0]][n[1]] = 2
		*(size)++
	}
	for _, n := range nodes {
		bfs(m, n, size)
	}
}

func isValid(m [][]int32, node []int32) bool {
	r := node[0]
	c := node[1]

	if r < 0 || r >= int32(len(m)) || c < 0 || c >= int32(len(m[1])) {
		return false
	}
	return true && m[r][c] == 1
}

func getAdjacet(m [][]int32, node []int32) [][]int32 {
	directions := [][]int32{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
	var nodes [][]int32
	for _, d := range directions {
		newNode := []int32{node[0] - d[0], node[1] - d[1]}
		if isValid(m, newNode) {
			nodes = append(nodes, newNode)
		}
	}
	return nodes
}
