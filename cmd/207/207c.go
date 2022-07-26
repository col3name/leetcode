package _07

//
//type Graph [][]int
//
//func canFinish(numCourses int, prerequisites [][]int) bool {
//	indegree := make([]int, numCourses)
//	adj := make(map[int][]int, 0)
//	var temp []int
//	for _, pre := range prerequisites {
//		key := pre[0]
//		val, ok := adj[key]
//		if ok {
//			temp = val
//		} else {
//			temp = make([]int, 0)
//		}
//		temp = append(temp, pre[1])
//		adj[key] = temp
//		indegree[pre[1]]++
//	}
//	queue := make([]int, 0)
//	for i := 0; i < numCourses; i++ {
//		if indegree[i] == 0 {
//			queue = append(queue, i)
//		}
//	}
//	count := 0
//	for len(queue) > 0 {
//		pre := queue[0]
//		queue = queue[1:]
//		val, ok := adj[pre]
//		if ok {
//			for _, v := range val {
//				indegree[v]--
//				if indegree[v] == 0 {
//					queue = append(queue, v)
//				}
//			}
//		}
//	}
//	return count == numCourses
//}
//
//func util(graph Graph, v int, visited []bool, stack []int) ([]int, bool) {
//	visited[v] = true
//	ok := false
//	for from := range graph[v] {
//		if !visited[from] {
//			stack, ok = util(graph, from, visited, stack)
//			if !ok {
//				return stack, false
//			}
//		}
//	}
//	stack = append(stack, v)
//	return stack, true
//}
