func orangesRotting(grid [][]int) int {
    height, width := len(grid), len(grid[0])
	rottenQueue := [][]int{}
	fresh := 0

	for row, rowv := range grid {
		for col, v := range rowv {
			if v == 1 {
				fresh++
			}
			if v == 2 {
				rottenQueue = append(rottenQueue, []int{row, col})
			} 
		} 
	}

	if fresh == 0 {
		return 0
	} 

	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	minutes := 0

	for fresh > 0 && len(rottenQueue) > 0 {
		countOfRottenFruitsOnLevel := len(rottenQueue)

		for i := 0; i < countOfRottenFruitsOnLevel; i++ {
			curRotten := rottenQueue[0]
			rottenQueue = rottenQueue[1:]

			row, col := curRotten[0], curRotten[1]

			for _, v := range dirs {
				newRow, newCol := row + v[0], col + v[1]

				if newRow >= 0 && newCol >=0 && newRow < height && newCol < width && grid[newRow][newCol] == 1 {
					grid[newRow][newCol] = 2
					fresh--
					rottenQueue = append(rottenQueue, []int{newRow, newCol})
				}

			}
		}
		minutes++
	}

	if fresh > 0 {
		return - 1
	}

	return minutes
}
