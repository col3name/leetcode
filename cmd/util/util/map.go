package util

import "fmt"

type Data struct {
	Value int `json:"value"`
}

func MapP(data []Data, fn func(item Data) interface{}) <-chan interface{} {
	out := make(chan interface{}, 10000)
	go func() {
		for _, item := range data {
			out <- fn(item)
		}
		close(out)
	}()
	return out
}

func FilterP(in <-chan interface{}, fn func(item interface{}) bool) <-chan interface{} {
	out := make(chan interface{}, 10000)

	go func() {
		for item := range in {
			i := fn(item)
			if i {
				out <- item
			}
		}
		close(out)
	}()
	return out
}

func ReduceP(in <-chan interface{},
	accumulator interface{},
	fn func(item interface{},
		accumulator interface{}) interface{}) <-chan interface{} {

	out := make(chan interface{}, 1)
	go func() {
		for item := range in {
			accumulator = fn(item, accumulator)
		}
		fmt.Println("accum", accumulator)
		out <- accumulator
		close(out)
	}()

	return out
}

func Map(list []Data, fn func(item Data) interface{}) []interface{} {
	res := make([]interface{}, len(list))
	for i, item := range list {
		res[i] = fn(item)
	}
	return res
}

func Filter(list []interface{}, fn func(item interface{}) bool) []interface{} {
	res := make([]interface{}, 0)
	for _, item := range list {
		if ok := fn(item); ok {
			res = append(res, item)
		}
	}
	return res
}

func Reduce(list []interface{},
	accumulator interface{},
	fn func(item interface{},
		accumulator interface{}) interface{}) interface{} {

	for _, item := range list {
		accumulator = fn(item, accumulator)
	}
	return accumulator
}
