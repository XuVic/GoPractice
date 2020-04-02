package search

import "sort"

func numberOfPairs(k int, arr []int) int {
	sort.Ints(arr)
	i := 0
	j := 1
	count := 0

	for j < len(arr) {
		diff := arr[j] - arr[i]
		if diff == k {
			count++
			j++
		} else if diff < k {
			j++
		} else if diff > k {
			i++
		}
	}
	return count
}

func transform(k int32, arr []int32) (int, []int) {
	k2 := int(k)
	arr2 := make([]int, len(arr))
	for i, val := range arr {
		arr2[i] = int(val)
	}
	return k2, arr2
}
