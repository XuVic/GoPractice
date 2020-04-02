package dp

func numberOfPalindrome(s string) int {
	var c1 [26]int
	var c2 [26][26]int
	var c3 [26][26]int
	count := 0

	for i := 0; i < len(s); i++ {
		cInt := int(s[i]) - int('a')
		for j := 0; j < 26; j++ {
			count += c3[cInt][j] % 100000007
			c3[j][cInt] += c2[j][cInt] % 100000007
			c2[j][cInt] += c1[j] % 100000007
		}
		c1[cInt]++
	}
	return count % 100000007
}
