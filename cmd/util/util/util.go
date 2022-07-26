package util

import (
	"fmt"
	"time"
)

func Times(fn func() int) {
	start := time.Now()
	i := fn()
	fmt.Println(time.Since(start), i)
}

func Gen() []Data {
	list := []Data{
		{Value: 0},
		{Value: 1},
		{Value: 2},
		{Value: 3},
		{Value: 4},
	}
	for i := 56; i < 100000; i++ {
		list = append(list, Data{Value: i})
	}
	return list
}
