func maxArea(heights []int) int {
	mx := 0
	l, r := 0, len(heights) - 1

	for l < r {
		water := min(heights[l], heights[r]) * (r - l)

		if water > mx {
			mx = water
		}

		if heights[l] > heights[r] {
			r--
		} else {
			l++
		}
	}

	return  mx
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
