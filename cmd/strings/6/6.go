package _6

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	rows := make([]strings.Builder, 0)
	length := numRows
	if length < len(s) {
		length = len(s)
	}

	for i := 0; i < length; i++ {
		rows = append(rows, strings.Builder{})
	}
	currentRow := 0
	goingDown := false
	for _, char := range []byte(s) {
		rows[currentRow].Write([]byte{char})
		if currentRow == 0 || currentRow == numRows-1 {
			goingDown = !goingDown
		}
		if goingDown {
			currentRow += 1
		} else {
			currentRow += -1
		}
	}
	var result strings.Builder
	{
	}
	for _, row := range rows {
		result.WriteString(row.String())
	}
	return result.String()
}
