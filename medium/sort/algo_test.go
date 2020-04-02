package sort

import "testing"

var x = [][]string{{"0", "a"}, {"1", "b"}, {"0", "c"}, {"1", "d"}}

func TestAlgo(t *testing.T) {
	t.Log(makeMap(x))
	countSort(x)
}
