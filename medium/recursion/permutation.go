package recursion

func permutation(slice []int) [][]int {
	if len(slice) == 1 {
		return [][]int{slice}
	}

	var res [][]int

	for i, val := range slice {
		var newSlice []int
		newSlice = append(newSlice, slice[:i]...)
		newSlice = append(newSlice, slice[i+1:]...)
		subSlices := permutation(newSlice)
		for _, subSlice := range subSlices {
			permu := []int{val}
			permu = append(permu, subSlice...)
			res = append(res, permu)
		}
	}
	return res
}
