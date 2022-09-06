package _1311

import "sort"

type movieStatistic struct {
	key   string
	count int
}

// 1311. Get Watched Videos by Your Friends

func watchedVideosByFriends(watchedVideos [][]string, friends [][]int, id int, level int) []string {
	queue := traverseToTargetLevel(friends, id, level)
	friendSet := getUniqueFriendsId(queue, id)
	frequency := calculateMovieFrequency(watchedVideos, friendSet)
	movieStatistics := buildMovieStatistic(frequency)
	return fillResult(movieStatistics)
}

func traverseToTargetLevel(friends [][]int, id int, targetLevel int) []int {
	currentLevel := 0
	visited := make([]bool, len(friends))

	q := make([]int, 0, 100)
	q = append(q, id)

	for true {
		size := len(q)
		for i := 0; i < size; i++ {
			friendId := q[0]
			q = q[1:]
			for _, frId := range friends[friendId] {
				if !visited[frId] {
					q = append(q, frId)
					visited[frId] = true
				}
			}
		}
		currentLevel++
		if targetLevel == currentLevel {
			break
		}
	}
	return q
}

func getUniqueFriendsId(q []int, id int) map[int]bool {
	friendSet := make(map[int]bool)

	for len(q) > 0 {
		friendId := q[0]
		q = q[1:]
		if friendId != id {
			if _, ok := friendSet[friendId]; !ok {
				friendSet[friendId] = true
			}
		}
	}
	return friendSet
}

func calculateMovieFrequency(watchedVideos [][]string, friendSet map[int]bool) map[string]int {
	frequency := make(map[string]int, len(watchedVideos))
	for friendId := range friendSet {
		for _, movie := range watchedVideos[friendId] {
			frequency[movie]++
		}
	}
	return frequency
}

func buildMovieStatistic(frequency map[string]int) []*movieStatistic {
	movieStatistics := make([]*movieStatistic, 0, len(frequency))
	for movie, freq := range frequency {
		movieStatistics = append(movieStatistics, &movieStatistic{key: movie, count: freq})
	}

	sort.Slice(movieStatistics, func(i, j int) bool {
		left := movieStatistics[i]
		right := movieStatistics[j]
		if left.count != right.count {
			return left.count < right.count
		}
		return left.key < right.key
	})

	return movieStatistics
}

func fillResult(tmp []*movieStatistic) []string {
	var result []string
	for _, movieStat := range tmp {
		result = append(result, movieStat.key)
	}
	return result
}
