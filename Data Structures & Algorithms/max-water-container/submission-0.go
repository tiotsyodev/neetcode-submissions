
func maxArea(heights []int) int {
	maxWater := 0

	l := 0
	r := len(heights) - 1

	for r > l {
		water := (r-l) * min(heights[r], heights[l])

		if water > maxWater {
			maxWater = water
		}

		if heights[l] > heights[r] {
			r--
		} else {
			l++
		}
	}

	return maxWater
}