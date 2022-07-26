package _3

type val struct {
	val     int
	nextOne string
	nextTwo string
}

func romanToInt2(s string) int {
	pair := map[string]val{
		"I": {1, "V", "X"},
		"V": {5, "", ""},
		"X": {10, "L", "C"},
		"L": {50, "", ""},
		"C": {100, "D", "M"},
		"D": {500, "", ""},
		"M": {1000, "", ""},
	}

	res := 0

	for i := 0; i < len(s); i++ {
		cur := string(s[i])
		Cur := pair[cur]
		if i < len(s)-1 {
			next := string(s[i+1])
			Next := pair[next]
			if Cur.nextOne == next || Cur.nextTwo == next {
				res += Next.val - Cur.val
				i += 1
			} else {
				res += Cur.val
			}
		} else {
			res += pair[cur].val
		}
	}
	return res
}

func romanToInt(s string) int {
	pair := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	res := 0

	for i := 0; i < len(s); i++ {
		cur := string(s[i])
		if (i < len(s)-1) && pair[cur] < pair[string(s[i+1])] {
			res -= pair[string(s[i])]
		} else {
			res += pair[string(s[i])]
		}
	}
	return res
}
