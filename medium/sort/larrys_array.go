package sort

import (
	"fmt"
	"reflect"
)

var A = []int32{1, 3, 4, 2}

func larrysArray(A []int32) string {
	swapper := reflect.Swapper(A)

	for i := len(A) - 1; i >= 2; i-- {
		for j := 2; j <= i; j++ {
			swap(swapper, A, j)
		}
	}
	fmt.Println(A)
	if A[0] < A[1] {
		return "YES"
	}
	return "NO"
}

func swap(swapper func(i, j int), A []int32, index int) {
	bigI := biggest(A, index)

	if bigI == index-1 {
		swapper(index, index-1)
		swapper(index-1, index-2)
	} else if bigI == index-2 {
		swapper(index, index-1)
		swapper(index-2, index)
	}

}

func biggest(A []int32, index int) int {
	if A[index] > A[index-1] {
		if A[index] > A[index-2] {
			return index
		}
		return index - 2
	}
	if A[index-1] > A[index-2] {
		return index - 1
	}
	return index - 2
}
