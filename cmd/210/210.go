package _10

var visited []bool
var edges [][]int

func dfs(course int, path []bool, stack []int) ([]int, bool) {
	if visited[course] {
		return nil, true
	}
	visited[course] = true
	path[course] = true
	ok := false
	for i := 0; i < len(edges[course]); i++ {
		stack, ok = dfs(edges[course][i], path, nil)
		if path[edges[course][i]] || !ok {
			return nil, false
		}
	}
	path[course] = false
	stack = append(stack, course)
	return stack, true
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	edges = make([][]int, numCourses)
	for _, pre := range prerequisites {
		key := pre[0]
		val := edges[key]
		if val == nil {
			val = make([]int, 0)
		}
		edges[key] = append(val, pre[1])
	}
	visited = make([]bool, numCourses)
	path := make([]bool, numCourses)
	stack := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		visited[i] = false
		path[i] = false
	}

	ok := false
	for i := 0; i < numCourses; i++ {
		stack, ok = dfs(i, path, stack)
		if !visited[i] && !ok {
			return []int{}
		}
	}
	return stack
}
