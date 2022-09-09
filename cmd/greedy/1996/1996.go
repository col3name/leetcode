package _1996

import "sort"

// 1996. The Number of Weak Characters in the Game
func numberOfWeakCharacters(properties [][]int) int {
	if len(properties) < 2 {
		return 0
	}
	sort.Slice(properties, func(i, j int) bool {
		if properties[i][0] == properties[j][0] {
			return properties[i][1] > properties[j][1]
		}
		return properties[i][0] < properties[j][0]
	})
	result := 0
	maxDef := 0
	for i := len(properties) - 1; i >= 0; i-- {
		if properties[i][1] < maxDef {
			result++
		} else {
			maxDef = properties[i][1]
		}
	}
	return result
}
