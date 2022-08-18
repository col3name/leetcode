package _51

// 51. N-Queens

type HashSet map[int]struct{}

func (s *HashSet) Contains(key int) bool {
	_, ok := (*s)[key]
	return ok
}

func (s *HashSet) Add(value int) {
	if _, ok := (*s)[value]; !ok {
		(*s)[value] = struct{}{}
	}
}
func (s *HashSet) Remove(value int) {
	delete(*s, value)
}

func solveNQueens(size int) [][]string {
	col := &HashSet{}
	diag1 := &HashSet{}
	diag2 := &HashSet{}
	board := make([][]string, size)
	for i := range board {
		board[i] = make([]string, size)
		for j := 0; j < size; j++ {
			board[i][j] = "."
		}
	}

	var result [][]string
	var helper func(row int, list []string)
	helper = func(row int, list []string) {
		if row == size {
			tmp := make([]string, len(list))
			for i, line := range list {
				tmp[i] = line
			}
			result = append(result, tmp)
			return
		}

		for i := 0; i < size; i++ {
			if col.Contains(i) || diag1.Contains(row+i) || diag2.Contains(row-i) {
				continue
			}

			list = append(list, fillRow(i, size))
			col.Add(i)
			diag1.Add(row + i)
			diag2.Add(row - i)

			helper(row+1, list)

			list = list[:len(list)-1]
			col.Remove(i)
			diag1.Remove(row + i)
			diag2.Remove(row - i)
		}
	}

	helper(0, []string{})
	return result
}

func fillRow(column, size int) string {
	arr := make([]byte, size)
	for j := 0; j < size; j++ {
		arr[j] = '.'
	}
	arr[column] = 'Q'
	return string(arr)
}
