func lengthOfLongestSubstring(s string) int {
	l := 0
	resmap := map[rune]int{}
	strMass := []rune(s)
	count := 0

	for r := 0; r < len(strMass); r++ {
		v, ok := resmap[strMass[r]]

		if ok && v >= l {
			l = v + 1	
		}

		resmap[strMass[r]] = r

		if r-l+1 > count {
			count = r - l + 1
		}
	}
	return count
}