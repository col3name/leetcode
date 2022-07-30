package _49

import "sort"

func groupAnagrams(strs []string) [][]string {
	hash := make(map[string][]string, 0)

	for _, str := range strs {
		ch := []byte(str)
		sort.Slice(ch, func(i, j int) bool {
			return ch[i] < ch[j]
		})
		s := string(ch)
		_, ok := hash[s]
		if ok {
			hash[s] = append(hash[s], str)
		} else {
			hash[s] = make([]string, 0)
			hash[s] = append(hash[s], str)
		}
	}
	result := make([][]string, 0)

	for _, value := range hash {
		result = append(result, value)
	}
	return result
}
