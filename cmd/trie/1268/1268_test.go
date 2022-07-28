package _1268

import (
	"fmt"
	"testing"
)

func TestName3(t *testing.T) {
	arr := []string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}
	products := suggestedProducts(arr, "mouse")
	fmt.Println(products)
}
