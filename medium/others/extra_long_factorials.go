package others

import (
	"strconv"
	"strings"
)

func extraLongFact(n int) string {
	res := "1"

	for i := 2; i <= n; i++ {
		res = stringMul(res, i)
	}
	return res
}

func stringMul(s1 string, i int) string {
	res := make([]string, len(s1)+1)
	ext := 0

	for n := len(s1) - 1; n >= 0; n-- {
		m, _ := strconv.Atoi(string(s1[n]))
		m = m*i + ext
		d := int(m % 10)
		ext = m / 10
		res[n] = strconv.Itoa(d)
	}

	for ext != 0 {
		d := int(ext % 10)
		ext = ext / 10
		res = append([]string{strconv.Itoa(d)}, res...)
	}

	return strings.Join(res, "")
}
