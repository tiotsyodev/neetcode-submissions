func findOrder(numCourses int, prerequisites [][]int) []int {
    graph := make([][]int, numCourses)

	for _, prereq := range prerequisites {
		course := prereq[0]
		pre := prereq[1]

		graph[pre] = append(graph[pre], course)
	}

	res := []int{}
	state := make([]int, numCourses)
	var dfs func(i int) bool
	dfs = func(i int) bool {
		if state[i] == 1 {
			return false
		}

		if state[i] == 2 {
			return true
		}

		state[i] = 1

		for _, v := range graph[i] {
			if !dfs(v) {
				return false
			}
		}

		state[i] = 2
		res = append(res, i)
		return true
	}

	for i := 0; i < numCourses; i++ {
		if state[i] == 0 {
			if !dfs(i) {
				return []int{}
			}
		}
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
        res[i], res[j] = res[j], res[i]
    }

	return res
}
