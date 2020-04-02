package others

import (
	"math"
	"strings"
)

func encrypt(s string) string {
	if len(s) == 1 {
		return s
	}

	s = strings.TrimSpace(s)
	sMatrix := makeMatrix(s)
	var res strings.Builder

	r := len(sMatrix)
	c := len(sMatrix[1])

	for i := 0; i < c; i++ {
		for j := 0; j < r; j++ {
			res.WriteString(sMatrix[j][i])
		}
		res.WriteString(" ")
	}
	return res.String()
}

func makeMatrix(s string) [][]string {
	sL := len(s)
	lSqrt := math.Sqrt(float64(sL))
	r := int(math.Floor(lSqrt))
	c := int(math.Ceil(lSqrt))
	if r*c < sL {
		r++
	}
	res := make([][]string, r)
	for i := range res {
		res[i] = make([]string, c)
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if (i*c)+j >= len(s) {
				break
			}
			res[i][j] = string(s[(i*c)+j])
		}
	}
	return res
}
