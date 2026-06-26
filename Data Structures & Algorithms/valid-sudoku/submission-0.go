func isValidSudoku(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		row := map[int]bool{}
		col := map[int]bool{}


		for j := 0; j < len(board[i]); j++ {
			// row
			if board[i][j] != '.' {
				v, _ := strconv.Atoi(string(board[i][j]))
				if row[v] {
					return false
				}
				row[v] = true
			}
			// col
			if board[j][i] != '.' {
				v, _ := strconv.Atoi(string(board[j][i]))
				if col[v] {
					return false
				}
				col[v] = true
			}
		}
	}

	for box := 0; box < 9; box++ {
		seen := make(map[byte]bool)
		startRow := (box / 3) * 3
		startCol := (box % 3) * 3
		
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if board[startRow+i][startCol+j] == '.' {
					continue
				}

				if seen[board[startRow+i][startCol+j]] {
					return false
				}
				
				seen[board[startRow+i][startCol+j]] = true
			}
		}
	}



	return true

}