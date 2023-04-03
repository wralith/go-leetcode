package leetcode

func canFinish(numCourses int, prerequisites [][]int) bool {
	g := map[int][]int{}

	for _, p := range prerequisites {
		g[p[1]] = append(g[p[1]], p[0])
	}

	visitSet := map[int]bool{}

	var dfs func(course int) bool
	dfs = func(course int) bool {
		// Loop
		if _, ok := visitSet[course]; ok {
			return false
		}

		// No prerequisite
		if len(g[course]) == 0 {
			return true
		}

		visitSet[course] = true // dfs on it!

		for _, p := range g[course] {
			if !dfs(p) {
				return false // if a course is impossible to complete
			}
		}

		delete(visitSet, course) // can be taken! - will be fall into no prerequisite from now on
		g[course] = []int{}

		return true
	}

	// All the courses event they are not connected in graph
	for i := 0; i < numCourses; i++ {
		if !dfs(i) {
			return false
		}
	}

	return true
}
