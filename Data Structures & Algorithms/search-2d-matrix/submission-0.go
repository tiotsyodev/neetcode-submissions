func searchMatrix(matrix [][]int, target int) bool {

	for _, v := range matrix {
		if v[0] <= target && v[len(v) - 1] >= target {

			l := 0
			r := len(v) - 1

			for r >= l {
				pivot := l + (r - l) / 2
				if v[pivot] > target {
					r = pivot - 1
				} else if v[pivot] < target {
					l = pivot + 1
				} else {
					return true
				}
			}
		}
	}

	return false
}
