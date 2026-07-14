func islandsAndTreasure(grid [][]int) {
	queue := [][]int{}
	dirs := [][]int{{0, 1}, {1, 0}, {-1, 0}, { 0, -1}}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				queue = append(queue, []int{i, j})
			}
		}
	}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		row, col := cur[0], cur[1]

		for _, v := range dirs {
			nrow, ncol := row + v[0], col + v[1]

			if nrow >= 0 && ncol >= 0 && nrow < len(grid) && ncol < len(grid[0]) && grid[nrow][ncol] == 2147483647 {
				grid[nrow][ncol] = grid[row][col] + 1
				queue = append(queue, []int{nrow, ncol})
			}
		}

	}
}
