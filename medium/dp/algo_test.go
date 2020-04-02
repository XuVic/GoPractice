package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testStr string = ""

var n int = 10
var coins = []int{2, 5, 3, 6}

func TestAlgo(t *testing.T) {
	assert.Equal(t, 4, numberOfChange(n, coins))
}

func TestOthers(t *testing.T) {
	// arr := initDp(n, coins)
	// t.Log(arr)
}
