package others

func countSubset(k int32, s []int32) int32 {
	reMap := normalize(k, s)
	var max int32 = 0
	for i := int32(1); i <= int32(k/2); i++ {
		if i == k-i && reMap[i] > 0 {
			max++
		} else if reMap[i] > reMap[k-i] {
			max += reMap[i]
		} else {
			max += reMap[k-i]
		}
	}
	if reMap[0] > 0 {
		max++
	}
	return max
}

func normalize(k int32, s []int32) map[int32]int32 {
	res := make(map[int32]int32, 1)
	for i := range s {
		d := int32(s[i] % k)
		_, ok := res[d]
		if !ok {
			res[d] = 1
		} else {
			res[d]++
		}
	}
	return res
}
