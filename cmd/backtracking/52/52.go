package _52

// 52. N-Queens II

type HashSet map[int]struct{}

func (h *HashSet) Add(value int) {
	(*h)[value] = struct{}{}
}
func (h *HashSet) Delete(value int) {
	delete(*h, value)
}
func (h *HashSet) Contains(value int) bool {
	_, ok := (*h)[value]
	return ok
}

func totalNQueens(size int) int {
	col := &HashSet{}
	diag1 := &HashSet{}
	diag2 := &HashSet{}

	var result [][]string
	var helper func(row int, list []string)
	helper = func(row int, list []string) {
		if row == size {
			tmp := make([]string, size)
			for i, line := range list {
				tmp[i] = line
			}
			result = append(result, tmp)
			return
		}
		for column := 0; column < size; column++ {
			if col.Contains(column) || diag1.Contains(row+column) || diag2.Contains(row-column) {
				continue
			}

			col.Add(column)
			diag1.Add(row + column)
			diag2.Add(row - column)

			helper(row+1, append(list, fillLine(column, size)))

			col.Delete(column)
			diag1.Delete(row + column)
			diag2.Delete(row - column)
		}
	}

	helper(0, []string{})

	return len(result)
}

func fillLine(column, size int) string {
	arr := make([]byte, size)

	for j := range arr {
		arr[j] = '.'
	}
	arr[column] = 'Q'

	return string(arr)
}
