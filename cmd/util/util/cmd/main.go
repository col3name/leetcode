package main

import (
	"fmt"
	"github.com/col3name/algo/cmd/util/util"
	"sort"
	"sync/atomic"
)

func main() {
	list := util.Gen()
	fnMap := func(item util.Data) interface{} {
		return item.Value * 3
	}

	fnFilterOne := func(item interface{}) bool {
		switch item.(type) {
		case int:
			return item.(int)%2 == 0
		default:
			return false
		}
	}

	fnFilterTwo := func(item interface{}) bool {
		switch item.(type) {
		case int:
			return item.(int) > 6
		default:
			return false
		}
	}

	reduceFn := func(item interface{}, accumulator interface{}) interface{} {
		switch item.(type) {
		case int:
			return accumulator.(int) + item.(int)
		default:
			return accumulator
		}
	}
	var i atomic.Int64
	i.Add(1)
	fmt.Println(i.Load())


}


util.Times(func () int {
	out := util.ReduceP(util.FilterP(util.FilterP(util.MapP(list, fnMap), fnFilterOne), fnFilterTwo),
		0,
		reduceFn)
	res := <-out
	return res.(int)
})
}
