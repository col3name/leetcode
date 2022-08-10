package _844

func backspaceCompare(s string, t string) bool {
	return intoEmptyText(s) == intoEmptyText(t)
}

func intoEmptyText(x string) string {
	l := len(x)
	tmp := make([]byte, 0, l)
	for i := 0; i < l; i++ {
		if x[i] != '#' {
			tmp = append(tmp, x[i])
		} else if len(tmp) != 0 {
			tmp = tmp[:len(tmp)-1]
		}
	}
	return string(tmp)
}
