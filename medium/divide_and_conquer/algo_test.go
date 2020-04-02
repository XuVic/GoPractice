package divide_and_conquer

import "testing"

var n int32 = 4
var k int32 = 2

func TestAlgo(t *testing.T) {
	t.Log(absolutePermutation(n, k))
	t.Log(absolutePermutation(3, 0))
}
