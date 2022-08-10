package _59

func generateMatrix(n int) [][]int {
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}
	count := 1
	for layer := 0; layer < (n+1)/2; layer++ {
		for i := layer; i < n-layer; i++ {
			result[layer][i] = count
			count++
		}
		for i := layer + 1; i < n-layer; i++ {
			result[i][n-layer-1] = count
			count++
		}
		for i := n - layer - 2; i >= layer; i-- {
			result[n-layer-1][i] = count
			count++
		}
		for i := n - layer - 2; i > layer; i-- {
			result[i][layer] = count
			count++
		}
	}
	return result
}
