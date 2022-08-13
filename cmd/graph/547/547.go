package _547

func findCircleNum(graph [][]int) int {
	visited := make(map[int]bool, 0)

	var dfs func(int)
	dfs = func(v int) {
		visited[v] = true
		for i, item := range graph[v] {
			if item == 1 && !visited[i] {
				dfs(i)
			}
		}
	}

	count := 0
	for i := 0; i < len(graph); i++ {
		if !visited[i] {
			dfs(i)
			count++
		}
	}

	return count
}
