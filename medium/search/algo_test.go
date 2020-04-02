package search

import "testing"

var text = []string{"1234567890",
	"0987654321",
	"1211111141",
	"1321111151",
	"2222222222"}
var pattern = []string{"876543",
	"111111",
	"211111"}

var testText = []string{"2234", "1122"}
var testPattern = []string{"1", "1"}

var text2 = []string{
	"111111111111111",
	"111111111111111",
	"111111011111111",
	"111111111111111",
	"111111111111111",
}
var pattern2 = []string{
	"11111",
	"11111",
	"11110",
}

func TestAlog(t *testing.T) {
	t.Log(gridSearch(text, pattern))
	t.Log(gridSearch(testText, testPattern))
	t.Log(horspoolSearch(text2[0], pattern2[0], len(pattern2[0])-1))
	t.Log(gridSearch(text2, pattern2))
}
