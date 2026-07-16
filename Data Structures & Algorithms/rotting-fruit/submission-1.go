type Point struct {
	R int
	C int
}

func orangesRotting(grid [][]int) int {
   qu := []Point{}
   fresh := 0
   for r, row := range grid {
		for c, col := range row {
			if col == 2 {
				qu = append(qu, Point{ R: r, C: c,})
			}

			if col == 1 {
				fresh++
			}
		}
   }

   dirs := [][]int{{0,1},{1,0},{-1,0},{0, -1}}
   minutes := 0

   for len(qu) > 0 && fresh > 0 {
		levelLen := len(qu)

		for i := 0; i < levelLen; i++ {
			r := qu[0].R
			c := qu[0].C

			qu = qu[1:]

			for _, v := range dirs {
				nr, nc := r + v[0], c + v[1] 

				if nr >= 0 && nc >= 0 && nr < len(grid) && nc < len(grid[0]) && grid[nr][nc] == 1 {
					grid[nr][nc] = 2
					fresh--
					qu = append(qu, Point{R: nr, C: nc})
				}

			}

		}

		minutes++
    }

	if fresh != 0 {
		return -1
	}

   return minutes
}
