package matrix

import (
	"fmt"
	"reflect"
	"strings"
)

var emaGrid = []string{
	"BGGBGB",
	"GGGGGG",
	"GGGGGB",
	"GGGGGG",
	"BGGGGG",
	"BGBBGB",
}

func getSet(a interface{}, b interface{}) []interface{} {
	set := make([]interface{}, 0)
	hash := make(map[interface{}]bool)
	av := reflect.ValueOf(a)
	bv := reflect.ValueOf(b)

	for i := 0; i < av.Len(); i++ {
		el := av.Index(i).Interface()
		hash[el] = true
	}

	for i := 0; i < bv.Len(); i++ {
		el := bv.Index(i).Interface()
		if _, found := hash[el]; found {
			set = append(set, el)
		}
	}

	return set
}

type Plus struct {
	Nodes []string
	Area  int32
}

func (p *Plus) Overlap(another *Plus) bool {
	set := getSet(p.Nodes, another.Nodes)
	return len(set) > 0
}

func (p *Plus) Add(node string) {
	p.Nodes = append(p.Nodes, node)
	p.Area++
}

func (p *Plus) Clone() *Plus {
	return &Plus{Nodes: p.Nodes, Area: p.Area}
}

func newGrid(grid []string) [][]string {
	newGrid := make([][]string, len(grid))

	for i, str := range grid {
		newGrid[i] = strings.Split(str, "")
	}

	return newGrid
}

func twoPluses(grid []string) int32 {
	var pluses []*Plus
	for r, rows := range grid {
		for c := range rows {
			plus := findPluse(grid, r, c)
			pluses = append(pluses, plus...)
		}
	}

	var area int32 = 0

	for i := 0; i < len(pluses); i++ {
		for j := i + 1; j < len(pluses); j++ {
			if !pluses[i].Overlap(pluses[j]) {
				a := pluses[i].Area * pluses[j].Area
				if a > area {
					area = a
				}
			}
		}
	}
	return area
}

func findPluse(grid []string, r, c int) []*Plus {
	plus := &Plus{Nodes: []string{toStr(r, c)}, Area: 1}

	pluses := []*Plus{plus}

	valid := func(r, c int) bool {
		if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) || grid[r][c] == 'B' {
			return false
		}
		return true
	}

	i := 1
	for true {
		directions := [][]int{{i, 0}, {-1 * i, 0}, {0, i}, {0, -1 * i}}
		var validNodes []string

		for _, d := range directions {
			newR := r + d[0]
			newC := c + d[1]
			if valid(newR, newC) {
				validNodes = append(validNodes, toStr(newR, newC))
			}
		}

		if len(validNodes) < 4 {
			break
		}

		for _, node := range validNodes {
			plus.Add(node)
		}

		pluses = append(pluses, plus.Clone())

		i++
	}
	return pluses
}

func toStr(r, c int) string {
	return fmt.Sprintf("%d,%d", r, c)
}
