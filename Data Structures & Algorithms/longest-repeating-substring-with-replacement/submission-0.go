func characterReplacement(s string, k int) int {
	l := 0
	ms := []rune(s)
	count := make([]int, 26)
	maxFreq := 0
	maxLength := 0

	for r, ch := range ms {

		idx := ch - 'A'
		count[idx]++
		maxFreq = max(count[idx], maxFreq)

		currWindow := r - l + 1

		if currWindow - maxFreq > k {
			count[ms[l] - 'A']--
			l++
		}

		if r - l + 1 > maxLength {
			maxLength = r - l + 1
		}
	}

	return maxLength
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}