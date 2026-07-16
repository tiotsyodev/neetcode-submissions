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
		r, c := queue[0][0], queue[0][1]
		queue =queue[1:]

		for _, v := range dirs {
			nr, nc := r + v[0], c + v[1] 
			if nr < 0 || nc < 0 || nr >= len(grid) || nc >= len(grid[0]) || grid[nr][nc] != 2147483647 {
				continue
			} 

			grid[nr][nc] = grid[r][c] + 1
			queue = append(queue, []int{nr, nc})
		}
		
	}
}
