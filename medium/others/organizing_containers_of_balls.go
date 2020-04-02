package others

func numberOf(container [][]int32) (numberOfC, numberOfB []int32) {
	numberOfC = make([]int32, len(container))
	numberOfB = make([]int32, len(container))
	for i := range container {
		for j := range container {
			numberOfC[i] += container[i][j]
			numberOfB[j] += container[i][j]
		}
	}
	return
}

func organize(container [][]int32) string {
	numberOfC, numberOfB := numberOf(container)

	res := "Possible"

	for i := range container {
		j := i
		for j < len(container) {
			if numberOfC[i] == numberOfB[j] {
				temp := numberOfB[j]
				numberOfB[j] = numberOfB[i]
				numberOfB[i] = temp
				break
			}
			j++
		}
		if j == len(container) {
			res = "Impossible"
		}
	}
	return res
}
