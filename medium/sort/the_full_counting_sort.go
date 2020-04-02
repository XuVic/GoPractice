package sort

import (
	"fmt"
	"strconv"
	"strings"
)

func countSort(arr [][]string) {
	resMap := makeMap(arr)

	var res strings.Builder

	for i := 0; i <= 99; i++ {
		str, ok := resMap[i]
		if ok {
			res.WriteString(strings.Join(str, " ") + " ")
		}
	}

	fmt.Println(strings.TrimRight(res.String(), " "))
}

func makeMap(arr [][]string) map[int][]string {
	resMap := make(map[int][]string)

	for j, val := range arr {
		str := val[1]
		if j < len(arr)/2 {
			str = "-"
		}

		i, _ := strconv.Atoi(val[0])
		if _, ok := resMap[i]; !ok {
			resMap[i] = []string{str}
		} else {
			resMap[i] = append(resMap[i], str)
		}
	}
	return resMap
}
