func canFinish(numCourses int, prerequisites [][]int) bool {
    graph := make([][]int, numCourses)
	for _, prereq := range prerequisites { 
		course := prereq[0]
		prereqCourse := prereq[1]
		graph[course] = append(graph[course], prereqCourse)
	}

	state := make([]int, numCourses)

	var dfs func(i int) bool
	dfs =  func(i int) bool {
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
		return true
	}

	for i := 0; i < numCourses; i++ {
		if state[i] == 0 {
			if !dfs(i) {
				return false
			}
		}
	}

	return true
}
