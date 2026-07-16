func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	for _, v := range prerequisites {
		course := v[0]
		pre := v[1]
		graph[course] = append(graph[course], pre)
	}

	state := make([]int, numCourses)
	var dfs func(course int) bool
	dfs = func(course int) bool {
		if state[course] == 1 {
			return false
		}

		if state[course] == 2 {
			return true
		}

		state[course] = 1

		for _, i := range graph[course] {
			if !dfs(i) {
				return false
			}
		}

		state[course] = 2
		return true

	}

	for i := 0; i < numCourses; i++ {
			if !dfs(i) {
				return false
			}
	}

	return true

	
}
