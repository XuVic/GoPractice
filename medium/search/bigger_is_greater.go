package search

import (
	"strings"
)

func solution(w string) string {
	nums := make([]rune, len(w))
	for i, char := range w {
		nums[i] = char
	}
	sortedNums := []rune{nums[len(nums)-1]}

	flag := false
	for i := len(nums) - 2; i >= 0; i-- {
		j, check := findSmallestBigger(sortedNums, nums[i])
		if check {
			nums = swap(sortedNums, nums, j, i)
			flag = true
			break
		} else {
			sortedNums = insertSort(sortedNums, nums[i])
		}
	}
	if flag {
		var res strings.Builder
		for _, charRune := range nums {
			res.WriteString(string(charRune))
		}
		if res.String() == w {
			return "no answer"
		}
		return res.String()
	}
	return "no answer"
}

func findSmallestBigger(sortedNums []rune, val rune) (int, bool) {
	for i, num := range sortedNums {
		if num > val {
			return i, true
		}
	}
	return 0, false
}

func swap(sortedNums []rune, nums []rune, j, i int) []rune {
	temp := nums[i]
	nums[i] = sortedNums[j]
	sortedNums = append(sortedNums[:j], sortedNums[j+1:]...)
	sortedNums = insertSort(sortedNums, temp)
	copy(nums[i+1:], sortedNums)
	return nums
}

func insertSort(sortedNums []rune, val rune) []rune {
	if len(sortedNums) == 0 {
		return []rune{val}
	}

	i := binarySearch(sortedNums, val)
	return Insert(sortedNums, val, i)
}

func binarySearch(sortedNums []rune, val rune) int {
	left := 0
	right := len(sortedNums) - 1
	mid := int((right + left) / 2)

	for right-left > 1 {
		if val == sortedNums[mid] {
			return mid
		} else if val > sortedNums[mid] {
			left = mid
		} else {
			right = mid
		}
		mid = int((right + left) / 2)
	}

	if val > sortedNums[right] {
		return right + 1
	} else if val > sortedNums[left] {
		return right
	} else if left == 0 {
		return 0
	}
	return left - 1
}

func Insert(slice []rune, val rune, index int) []rune {
	n := len(slice) - 1
	slice = append(slice, 0)
	copy(slice[index+1:], slice[index:n+1])
	slice[index] = val
	return slice
}
