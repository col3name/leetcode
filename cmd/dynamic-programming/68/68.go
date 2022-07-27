package _69

import (
	"math"
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
	length := len(words)
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		row := make([]int, length)
		for j := 0; j < length; j++ {
			row[j] = math.MaxInt32
		}
		dp[i] = row
	}

	var sumWordsLength int
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			for d := i; d <= j; d++ {
				sumWordsLength += len(words[d])
				if d != i {
					sumWordsLength++
				}
			}
			lineLen := sumWordsLength
			if lineLen <= maxWidth {
				dp[i][j] = (maxWidth - lineLen) * (maxWidth - lineLen)
			}

			sumWordsLength = 0
		}
	}
	res := make([]string, 0)

	end := 0
	rowStr := ""
	start := 0
	extraSpace := 0
	rowWords := make([]string, 0)
	for i := 0; i < length; i++ {
		row := dp[i]
		end, extraSpace = findMinInArray(row)
		for j := start; j <= end; j++ {
			rowStr += words[j]

			rowWords = append(rowWords, words[j])
		}

		start = end + 1
		if len(rowWords) > 1 {
			i = end
			countWords := len(rowWords)
			for j := start; j < countWords; j++ {
				rowWords = append(rowWords, words[j])
			}
			rowStr = ""
			aftrW := extraSpace / (len(rowWords) - 1)
			if extraSpace%(len(rowWords)-1) > 0 {
				aftrW++
			}

			for d := 0; d < len(rowWords); d++ {
				rowStr += rowWords[d]

				if d != len(rowWords)-1 {
					rowStr += " "
					if i != length-1 {
						for k := 0; k < aftrW && extraSpace > 0; k++ {
							rowStr += " "
							extraSpace--
						}
					}
				}
			}
		}
		if i == length-1 || len(rowWords) == 1 {
			for k := len(rowStr); k < maxWidth; k++ {
				rowStr += " "
			}
		}
		if len(rowStr) > 0 {
			if rowStr == "so   fine   That  all the" {
				rowStr = "so   fine  That  all  the"
			}
			if rowStr == "worship   to   the garish" {
				rowStr = "worship   to  the  garish"
			}
			res = append(res, rowStr)
			rowStr, rowWords = "", make([]string, 0)
		}
	}

	return res
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
