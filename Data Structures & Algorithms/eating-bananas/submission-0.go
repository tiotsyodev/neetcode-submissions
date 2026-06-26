func minEatingSpeed(piles []int, h int) int {
	r := 0

	for _, v := range piles {
		if v > r {
			r = v
		}
	}
	l := 1
	for l <= r {
		mid := l + (r - l) / 2
		hours := 0
		for i := 0; i < len(piles); i++ {
			hours += (piles[i] + mid - 1) / mid
		}

		if hours <= h {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return l
}
