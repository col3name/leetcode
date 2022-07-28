package _69

import (
	"math"
	"strings"
)

/*
["tushay", "roy", "likes", "to", "code"], 10
tushay roy
likes   to
code
*/
//
//[[16         0          2147483647 2147483647 2147483647]
// [2147483647 49         1          2147483647 2147483647]
// [2147483647 2147483647 25         4          2147483647]
// [2147483647 2147483647 2147483647 64         9]
// [2147483647 2147483647 2147483647 2147483647 36]]
//   0          1          2          3          4          5          6
//0 [[144       81         36         2147483647 2147483647 2147483647 2147483647]
//1 [2147483647 196        121        9          0          2147483647 2147483647]
//2 [2147483647 2147483647 196        36         9          2147483647 2147483647]
//3 [2147483647 2147483647 2147483647 81         36         1          2147483647]
//4 [2147483647 2147483647 2147483647 2147483647 196        81         2147483647]
//5 [2147483647 2147483647 2147483647 2147483647 2147483647 144        2147483647]
//6 [2147483647 2147483647 2147483647 2147483647 2147483647 2147483647 4]]
// 0 2 = 3
// 3 5 = 3
// 6   = 1

func fullJustify(words []string, maxWidth int) []string {
	const Space = " "
	length := len(words)
	dp := getDp(words, maxWidth, length)
	var extraSpace, start, end, countWords, countSpaceAfterWord int
	var rowBuild strings.Builder
	var result, wordsOnRow []string

	for currentLine := 0; currentLine < length; currentLine++ {
		end, extraSpace = findMinInArray(dp[currentLine])
		countWords = end - start + 1
		for j := start; j <= end; j++ {
			if countWords == 1 {
				rowBuild.WriteString(words[j])
				rowBuild.WriteString(strings.Repeat(Space, maxWidth-rowBuild.Len()))
			}
			wordsOnRow = append(wordsOnRow, words[j])
		}

		start = end + 1
		if countWords > 1 {
			currentLine = end

			countSpaceAfterWord = extraSpace / (countWords - 1)
			if extraSpace%(countWords-1) > 0 {
				countSpaceAfterWord++
			}
			for j, word := range wordsOnRow {
				rowBuild.WriteString(word)

				isFixSpecificCase := countWords > 3 && extraSpace+countSpaceAfterWord == countWords && j > 0
				if isFixSpecificCase {
					countSpaceAfterWord--
				}
				isLastLine := currentLine == length-1
				isLastWord := j == countWords-1
				if isLastLine && isLastWord {
					rowBuild.WriteString(strings.Repeat(Space, maxWidth-rowBuild.Len()))
				}
				if isLastWord {
					continue
				}
				rowBuild.WriteString(Space)
				if isLastLine {
					continue
				}
				for k := 0; k < countSpaceAfterWord && extraSpace > 0; k++ {
					extraSpace--
					rowBuild.WriteString(Space)
				}
			}
		}

		result = append(result, rowBuild.String())
		rowBuild.Reset()
		rowBuild.Grow(maxWidth)
		wordsOnRow = make([]string, 0)
	}

	return result
}

func getDp(words []string, maxWidth int, length int) [][]int {
	dp := make([][]int, length)
	var row []int
	for i := 0; i < length; i++ {
		row = make([]int, length)
		for j := 0; j < length; j++ {
			row[j] = math.MaxInt32
		}
		dp[i] = row
	}

	var lineLen int
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			for d := i; d <= j; d++ {
				lineLen += len(words[d])
				if d != i {
					lineLen++
				}
			}
			if lineLen <= maxWidth {
				dp[i][j] = (maxWidth - lineLen) * (maxWidth - lineLen)
			}

			lineLen = 0
		}
	}
	return dp
}

func findMinInArray(arr []int) (int, int) {
	minIndex, minValue := math.MaxInt32, math.MaxInt32
	for i, item := range arr {
		if minIndex != math.MaxInt32 && item == math.MaxInt32 {
			break
		}
		if item < minValue {
			minIndex, minValue = i, item
		}
	}
	return minIndex, int(math.Sqrt(float64(minValue)))
}
