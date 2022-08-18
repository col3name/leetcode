package bloom

import (
	"math"
)

type Filter struct {
	bitArray []bool
	arrSize  int
}

func NewFilter(size int) *Filter {
	return &Filter{bitArray: make([]bool, size), arrSize: size}
}

func (f *Filter) h1(s string) int {
	hash := 0
	for _, char := range []byte(s) {

		hash = hash + int(char)
		hash = hash % f.arrSize
	}

	return hash
}

func (f *Filter) h2(s string) int {
	hash := 1
	for i, char := range []byte(s) {
		hash = hash + int(math.Pow(float64(19), float64(i)))*int(char)
		hash = hash % f.arrSize
	}
	return hash % f.arrSize
}

func (f *Filter) h3(s string) int {
	hash := 7
	for _, char := range []byte(s) {
		hash = (hash*31 + int(char)) % f.arrSize
	}
	return hash % f.arrSize
}

func (f *Filter) h4(s string) int {
	hash := 3
	p := 7
	for i := 0; i < len(s); i++ {
		hash += hash*7 + int(s[0])*int(math.Pow(float64(p), float64(i)))
	}

	return hash % f.arrSize
}

func (f *Filter) Lookup(s string) bool {
	h1 := f.h1(s)
	h2 := f.h2(s)
	h3 := f.h3(s)
	h4 := f.h4(s)
	return f.bitArray[h1] && f.bitArray[h2] && f.bitArray[h3] && f.bitArray[h4]
}

func (f *Filter) Insert(s string) bool {
	if f.Lookup(s) {
		return false
	}
	f.bitArray[f.h1(s)] = true
	f.bitArray[f.h2(s)] = true
	f.bitArray[f.h3(s)] = true
	f.bitArray[f.h4(s)] = true
	return true
}
