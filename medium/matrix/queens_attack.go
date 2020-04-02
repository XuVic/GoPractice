package matrix

import (
	"fmt"
	"math"
)

func isObstacle(queen []int32, obstacle []int32) bool {
	qR := queen[0]
	qC := queen[1]
	oR := obstacle[0]
	oC := obstacle[1]

	res := qC == oC || qR == oR
	return res || math.Abs(float64(qR-oR)) == math.Abs(float64(qC-oC))
}

func initDistance(queen []int32, n int32) map[string]int32 {
	res := make(map[string]int32)
	qR := queen[0]
	qC := queen[1]

	res["up"] = n - qR
	res["down"] = qR - 1
	res["left"] = qC - 1
	res["right"] = n - qC
	res["ld"] = int32(math.Min(float64(qC-1), float64(qR-1)))
	res["lu"] = int32(math.Min(float64(qC-1), float64(n-qR)))
	res["rd"] = int32(math.Min(float64(n-qC), float64(qR-1)))
	res["ru"] = int32(math.Min(float64(n-qC), float64(n-qR)))
	return res
}

func numberOfAttacks(n int32, queen []int32, obstacles [][]int32) int32 {
	res := initDistance(queen, n)

	for _, o := range obstacles {
		if isObstacle(queen, o) {
			dir := getDirection(queen, o)
			if _, ok := res[dir]; !ok {
				res[dir] = distance(queen, o)
			} else {
				preD := res[dir]
				curD := distance(queen, o)
				fmt.Println(curD)
				fmt.Println(preD)
				if curD < preD {
					res[dir] = curD
					fmt.Println(res[dir])
				}
			}
		}
	}
	ans := int32(0)
	for _, v := range res {
		ans += v
	}
	fmt.Println(res)
	return ans
}

func distance(queen []int32, obstacle []int32) int32 {
	qR := queen[0]
	qC := queen[1]
	oR := obstacle[0]
	oC := obstacle[1]
	rD := math.Abs(float64(qR - oR))
	cD := math.Abs(float64(qC - oC))

	if qR == oR {
		return int32(cD) - 1
	} else if qC == oC {
		return int32(rD) - 1
	}

	return int32(math.Min(rD, cD)) - 1
}

func getDirection(queen []int32, obstacle []int32) string {
	qR := queen[0]
	qC := queen[1]
	oR := obstacle[0]
	oC := obstacle[1]

	if qR == oR {
		if qC > oC {
			return "left"
		}
		return "right"
	} else if qC == oC {
		if qR > oR {
			return "down"
		}
		return "up"
	}
	rD := qR - oR
	cD := qC - oC
	if rD > 0 {
		if cD > 0 {
			return "ld"
		}
		return "rd"
	}
	if cD > 0 {
		return "lu"
	}
	return "ru"
}
