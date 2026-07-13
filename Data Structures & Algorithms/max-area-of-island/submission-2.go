func maxAreaOfIsland(grid [][]int) int {
    rows := len(grid)
	cols := len(grid[0])

	res := 0
	cur := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {  
            	cur = 0
            	dfs(grid, &cur, rows, cols, i, j)
				if cur > res {
					res = cur
				}
       		}
		}	
	}

	return res
}	

func dfs(grid [][]int, cur *int, rows, cols, i, j int) {
	if i < 0 || j < 0 || i >= rows || j >= cols || grid[i][j] == 0 {
		return
	}

	*cur += 1

	grid[i][j] = 0

	dfs(grid, cur, rows, cols, i+1, j)
	dfs(grid, cur, rows, cols, i, j+1)
	dfs(grid, cur, rows, cols, i, j-1)
	dfs(grid, cur, rows, cols, i-1, j)
}
