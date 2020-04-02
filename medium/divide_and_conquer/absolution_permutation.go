package divide_and_conquer

func absolutePermutation(n int32, k int32) []int32 {
	nums := make(map[int32]bool)
	for i := int32(1); i <= n; i++ {
		nums[i] = false
	}

	res := make([]int32, n)

	for i := int32(1); i <= n; i++ {
		smaller := i - k
		bigger := k + i

		if take, ok := nums[smaller]; ok && take == false {
			res[i-1] = smaller
			nums[smaller] = true
		} else if take, ok := nums[bigger]; ok && take == false {
			res[i-1] = bigger
			nums[bigger] = true
		} else {
			return []int32{-1}
		}
	}
	return res

}

func permu(i int32, k int32, nums map[int32]bool, n int32) []int32 {
	bigger := k + i
	smaller := i - k

	if take, ok := nums[smaller]; ok && take == false {
		nums[smaller] = true
		subSlice := permu(i+1, k, nums, n-1)
		if len(subSlice) == int(n)-1 {
			res := []int32{smaller}
			return append(res, subSlice...)
		}
		nums[smaller] = false
	}

	if take, ok := nums[bigger]; ok && take == false {
		nums[bigger] = true
		subSlice := permu(i+1, k, nums, n-1)
		if len(subSlice) == int(n)-1 {
			res := []int32{bigger}
			return append(res, subSlice...)
		}
		nums[bigger] = false
	}
	return []int32{}
}
