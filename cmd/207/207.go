package _07

type Graph map[int][]int

var visited []bool
var edges [][]int

func canFinish(numCourses int, prerequisites [][]int) bool {
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
	for i := 0; i < numCourses; i++ {
		visited[i] = false
		path[i] = false
	}

	for i := 0; i < numCourses; i++ {
		if !visited[i] && !dfs(i, path) {
			return false
		}
	}
	return true
}

func dfs(course int, path []bool) bool {
	if visited[course] {
		return true
	}
	visited[course] = true
	path[course] = true
	for i := 0; i < len(edges[course]); i++ {
		if path[edges[course][i]] || !dfs(edges[course][i], path) {
			return false
		}
	}
	path[course] = false
	return true
}

func util(graph Graph, course int, visited []bool) bool {
	if visited[course] {
		return false
	}
	visited[course] = true
	for from := range graph[course] {
		if !util(graph, from, visited) {
			return false
		}
	}
	visited[course] = false
	return true
}
