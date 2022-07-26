package _857

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	colors := "abaca"
	edges := [][]int{{0, 1}, {0, 2}, {2, 3}, {3, 4}}
	value := largestPathValue(colors, edges)
	fmt.Println(value)
}
