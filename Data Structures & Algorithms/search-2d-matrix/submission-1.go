func searchMatrix(matrix [][]int, target int) bool {

	m := len(matrix)
	n := len(matrix[0])

	r := m * n - 1
	l := 0

	for l <= r {
		mid := l + (r - l) / 2
		row := mid / n
        col := mid % n
		val := matrix[row][col]

		if val > target {
			r = mid - 1
		} else if val < target {
			l = mid + 1
		} else {
			return true
		}

	}

	return false

}
