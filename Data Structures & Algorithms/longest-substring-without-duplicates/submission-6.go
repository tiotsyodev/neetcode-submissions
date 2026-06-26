func lengthOfLongestSubstring(s string) int {
	l := 0
	resmap := map[rune]int{}
	ms := []rune(s)
	count := 0

	for r, v := range ms {

		idx, ok := resmap[v]

		if ok && idx >= l {
			l = idx + 1
		}

		resmap[ms[r]] = r

		if r - l + 1 > count {
			count = r - l + 1
		}

	}

	return count
}