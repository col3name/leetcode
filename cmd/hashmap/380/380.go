package _380

import (
	"math/rand"
)

type RandomizedSet struct {
	dictionary map[int]int
	data       []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		dictionary: make(map[int]int, 0),
		data:       make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.dictionary[val]; ok {
		return false
	}
	this.dictionary[val] = len(this.data)
	this.data = append(this.data, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	itemIndex, ok := this.dictionary[val]
	if !ok {
		return false
	}
	lastIndex := len(this.data) - 1

	lastItem := this.data[lastIndex]

	//move last element to delete element position
	this.dictionary[lastItem] = itemIndex
	this.data[itemIndex] = lastItem

	delete(this.dictionary, val)
	this.data = this.data[:lastIndex]
	return true
}

func (this *RandomizedSet) GetRandom() int {
	max := len(this.data)
	if max == 0 {
		return this.data[0]
	}
	intn := rand.Intn(max)
	return this.data[intn]
}
