package _815
type Point struct {
	Node  int
	Depth int
}
type Queue []*Point
type MapSet map[int]struct{}
func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target { return 0}
	n, graph := len(routes), make([][]int, len(routes))
	for i := 0; i < n; i++ { graph[i] = make([]int, 0, 1000) }
    queue, seen,targets, foundSource, foundTarget := Queue(make([]*Point, 0, n)), MapSet(make(map[int]struct{}, n)), MapSet(make(map[int]struct{}, n)), false, false
	for i := 0; i < n; i++ {
		foundSource, foundTarget = Search(routes[i], source, target)
		if foundSource {seen[i] = struct{}{};queue = append(queue, &Point{i, 0})}
		if foundTarget { targets[i] = struct{}{}	}
		for j := i + 1; j < n; j++ {if intersect(routes[i], routes[j]) {graph[i] = append(graph[i], j);graph[j] = append(graph[j], i)}}
	}
	info, node, depth, ok := &Point{},0,0, false
	for len(queue) != 0 {
		info, queue = queue[0], queue[1:]
		node = info.Node
		depth = info.Depth
		if _, ok = targets[node]; ok { return depth + 1}
		for _, neighbour := range graph[node] {if _, ok = seen[neighbour]; ok {continue};seen[neighbour] = struct{}{}; queue = append(queue, &Point{neighbour, depth + 1})}
	}
	return -1
}

func Search(nums []int, targetOne, targetTwo int) (bool, bool) {
	foundOne, foundTwo :=false, false
	for _, num := range nums { if num == targetOne { foundOne = true};if num == targetTwo {foundTwo = true}; if foundTwo && foundOne {break}}
	return foundOne, foundTwo
}

func intersect(lhs []int, rhs []int) bool {
	hash := make(map[int]struct{})
	for _, item := range lhs { hash[item] = struct{}{}}
	ok := false
	for _, item := range rhs { if _, ok = hash[item]; ok {return true}}
	return false
}