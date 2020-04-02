package others

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, ok := keys[entry]; !ok {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func rankingAlice(ranks []int, alice []int) []int {
	res := make([]int, len(alice))

	rank := len(ranks) - 1

	for i, score := range alice {
		for rank > 0 && score >= ranks[rank] {
			rank--
		}

		aRank := rank + 1

		if ranks[rank] > score {
			aRank++
		}

		res[i] = aRank
	}

	return res
}
