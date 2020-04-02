package dp

import (
	"sort"
)

func initDp(n int, coins []int) [][]int {
	sort.Ints(coins)
	max := coins[len(coins)-1]
	matrix := make([][]int, n+1)
	for i := range matrix {
		matrix[i] = make([]int, max+1)
		for _, coin := range coins {
			if i == coin {
				matrix[i][coin]++
			}
		}
	}
	return matrix
}

func numberOfChange(n int, coins []int) int {
	dpMatrix := initDp(n, coins)

	for i := 2; i <= n; i++ {
		for _, coin := range coins {
			if coin >= i {
				break
			}
			dpMatrix[i][coin] += sumOfSlice(dpMatrix[i-coin], coin)
		}
	}
	return sumOfSlice(dpMatrix[n], 0)
}

func sumOfSlice(s []int, i int) (sum int) {
	for i < len(s) {
		sum += s[i]
		i++
	}
	return
}
