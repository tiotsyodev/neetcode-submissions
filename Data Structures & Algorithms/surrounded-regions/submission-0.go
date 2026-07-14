func solve(board [][]byte) {
    row, col := len(board), len(board[0])
	
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || c < 0 || r >= row || c >= col || board[r][c] != 'O' {
			return
		}

		board[r][c] = '*'

		dfs(r + 1, c)
		dfs(r, c + 1)
		dfs(r-1, c)
		dfs(r, c-1)
	}

	for i := 0; i < row; i++ {
		dfs(i, 0)
		dfs(i, col - 1)
	}

	for i := 0; i < col; i++ {
		dfs(0, i)
		dfs(row - 1, i)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if  board[i][j] == '*' {
				board[i][j] = 'O'
			}

		}
	}
	
}
