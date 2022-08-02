package _242

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := [26]int{}
	for _, r := range s {
		v := r - 'a'
		m[v] += 1
	}
	n := [26]int{}
	for _, r := range t {
		v := r - 'a'
		n[v] += 1
	}
	return m == n
}
