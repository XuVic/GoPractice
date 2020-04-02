package sort

import "fmt"

var sortA = []int32{1, 5, 4, 3, 2, 6}
var sortB = []int32{3, 1, 2}

func almostSorted(arr []int32) {
	if len(arr) == 2 {
		fmt.Println("yes")
		if arr[0] > arr[1] {
			fmt.Println("swap 1 2")
		}
	} else {
		desPoints := findCheckPoint(arr)
		res, str := checkCheckPoint(arr, desPoints)
		if res {
			fmt.Println("yes")
			fmt.Println(str)
		} else {
			fmt.Println("no")
		}
	}
}

func findCheckPoint(arr []int32) []int {
	preI := 0

	desPoints := make([]int, 0)

	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[preI] {
			if len(desPoints) == 0 {
				desPoints = append(desPoints, i)
			} else if preI == desPoints[len(desPoints)-1] {
				desPoints = append(desPoints, i)
			} else {
				desPoints = append(desPoints, -1)
				desPoints = append(desPoints, i)
			}
		}
		preI = i
	}
	return desPoints
}

func checkCheckPoint(arr []int32, desPoints []int) (bool, string) {
	lenD := len(desPoints)
	if lenD == 1 {
		if arr[desPoints[0]-1] < arr[desPoints[0]+1] && arr[desPoints[0]] > arr[desPoints[0]-2] {
			return true, fmt.Sprintf("swap %d %d", desPoints[0], desPoints[0]+1)
		}

		return false, ""
	}

	if contains(desPoints, -1) {
		if len(desPoints) > 3 {
			return false, ""
		}

		biggerI := desPoints[0] - 1
		smallerI := desPoints[2]

		if arr[biggerI] > arr[smallerI] {
			return true, fmt.Sprintf("swap %d %d", biggerI+1, smallerI+1)
		}
		return false, ""
	}

	leftI := desPoints[0] - 1
	rightI := desPoints[lenD-1]

	if arr[leftI] < arr[rightI+1] {
		if leftI == 0 || arr[rightI] > arr[leftI-1] {
			return true, fmt.Sprintf("reverse %d %d", leftI+1, rightI+1)
		}
	}

	return false, ""
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
