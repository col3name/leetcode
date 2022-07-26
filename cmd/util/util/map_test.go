package util

import (
	"fmt"
	"sort"
	"testing"
)

func TestMap(t *testing.T) {
	list := Gen()
	Times(func() int {
		list2 := Map(list, func(item Data) interface{} {
			return item.Value * 3
		})
		list3 := Filter(list2, func(item interface{}) bool {
			switch item.(type) {
			case int:
				return item.(int)%2 == 0
			default:
				return false
			}
		})
		sum := Reduce(list3, 0, func(item interface{}, accumulator interface{}) interface{} {
			switch item.(type) {
			case int:
				return accumulator.(int) + item.(int)
			default:
				return accumulator
			}
		})
		return sum.(int)
	})
}

func TestMapPipe(t *testing.T) {
	list := Gen()
	Times(func() int {
		fnMap := func(item Data) interface{} {
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

		out := ReduceP(FilterP(FilterP(MapP(list, fnMap), fnFilterOne), fnFilterTwo),
			0,
			reduceFn)
		res := <-out
		return res.(int)
	})
}

func TestFind(t *testing.T) {
	x := []int{1, 2, 3, 4, 5, 6, 7}
	target := 5
	sort.Slice(x, func(i, j int) bool {
		return x[i] < x[j]
	})
	fmt.Println(x)
	i, found := sort.Find(len(x), func(i int) int {
		if x[i] == target {
			return 0
		}
		if x[i] < target {
			return -1
		}
		return +1
	})
	if found {
		fmt.Printf("found %d at entry %d\n", target, i)
	} else {
		fmt.Printf("%d not found, would insert at %d", target, i)
	}
}
