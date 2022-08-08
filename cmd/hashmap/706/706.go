package _706

type MyHashMap struct {
	data []int
}

func Constructor() MyHashMap {
	return MyHashMap{
		data: make([]int, 1_000_001),
	}
}

func (this *MyHashMap) Put(key int, value int) {
	this.data[key] = value + 1
}

func (this *MyHashMap) Get(key int) int {
	return this.data[key] - 1
}

func (this *MyHashMap) Remove(key int) {
	this.data[key] = 0
}
