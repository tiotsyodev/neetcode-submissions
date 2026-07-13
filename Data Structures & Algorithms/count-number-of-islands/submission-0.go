func numIslands(grid [][]byte) int {
	res := 0
	rows := len(grid)
	cols := len(grid[0])
    
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				res++
				dfs(grid, i, j, rows, cols)
			}
		}
	}

	return res
}

func dfs(grid [][]byte, i, j, rows, cols int) {
	if i < 0 || j < 0 || i >= rows || j >= cols || grid[i][j] == '0'  {
		return
	}

	grid[i][j] = '0'

	dfs(grid, i+1, j, rows, cols)
	dfs(grid, i, j+1, rows, cols)
	dfs(grid, i-1, j, rows, cols)
	dfs(grid, i, j-1, rows, cols)

}
