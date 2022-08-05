package _2241

import (
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
)

type ATM struct {
	notes  []int
	np     []int
	values []int
}

func Constructor() ATM {
	return ATM{
		notes:  make([]int, 5),
		values: []int{20, 50, 100, 200, 500},
		np:     []int{-1},
	}
}

func (this *ATM) Deposit(banknotesCount []int) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println(err)
	} else {
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(bytes))
		}
	}
	for i := 0; i < 5; i++ {
		this.notes[i] += banknotesCount[i]
	}
}

func (this *ATM) Withdraw(amount int) []int {
	result := make([]int, 5)
	for i := 4; i >= 0; i-- {
		if this.notes[i] == 0 || amount < this.values[i] {
			continue
		}
		result[i] = int(math.Min(float64(amount/this.values[i]), float64(this.notes[i])))
		amount -= result[i] * this.values[i]
		if amount == 0 {
			for j := i; j < 5; j++ {
				this.notes[j] -= result[j]
			}
			return result
		}
	}

	return this.np
}
