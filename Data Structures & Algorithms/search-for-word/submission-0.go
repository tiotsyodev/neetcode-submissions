func exist(board [][]byte, word string) bool {
    bytes := []byte(word)
    lenv := len(board)
    lenh := len(board[0])

    visited := make([][]bool, lenv)
    for i := 0; i < len(visited); i++ {
        visited[i] = make([]bool, lenh)
    }

    var backtrack func(v, h, curIdx int) bool
    backtrack = func(v, h, curIdx int) bool {

        if curIdx == len(word) {
            return true
        }

        if v >= lenv || h >= lenh  || v < 0 || h < 0 {
            return false
        }

        if visited[v][h] {
            return false
        }

        if board[v][h] != bytes[curIdx] {
            return false
        }

        visited[v][h] = true

       found := backtrack(v + 1, h, curIdx + 1) ||
            backtrack(v - 1, h, curIdx + 1) ||
            backtrack(v, h + 1, curIdx + 1) ||
            backtrack(v, h - 1, curIdx + 1)

        visited[v][h] = false
        
        return found

    }

    for i := 0; i < lenv; i++ {
        for j := 0; j < lenh; j++ {
            if backtrack(i, j, 0) {
                return true
            }
        }
    }

    return false
}
