func minEatingSpeed(piles []int, h int) int {
	maxPile := 0

	for _, v := range piles {
		if v > maxPile {
			maxPile = v
		}
	}

	l := 1
	r := maxPile

	for r > l {
		mid := l + (r - l) / 2
		hours := 0

		for _, v := range piles {
			hours += (v + mid - 1) / mid
		}

		if hours <= h {
			r = mid 
		} else {
			l = mid + 1
		}
	}

	return l
}
