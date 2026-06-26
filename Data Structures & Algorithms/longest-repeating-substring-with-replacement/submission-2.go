func characterReplacement(s string, k int) int {
	l := 0
	count := make([]int, 26)
	maxFreq := 0
	counter := 0

	for r := 0; r < len(s); r++ {
		idx := int(s[r] - 'A')

		count[idx]++
		if maxFreq < count[idx] {
			maxFreq = count[idx]
		}

		if r - l + 1 - maxFreq > k {
			leftIdx := int(s[l] - 'A')
			l++
			count[leftIdx]--
		}

		if r - l + 1 > counter {
			counter = r - l + 1
		}
	}

	return counter
}
