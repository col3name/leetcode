package _659

// 659. Split Array into Consecutive Subsequences
func isPossible(nums []int) bool {
	mp := make(map[int]int, 0)
	for _, num := range nums {
		mp[num]++
	}

	var (
		count       int
		currentFreq int
		ok          bool
	)
	for _, num := range nums {
		maxFreq := mp[num]
		if maxFreq != 0 {
			count = 0

			currentFreq, ok = mp[num]
			for ok && currentFreq >= maxFreq {
				if currentFreq > maxFreq {
					maxFreq = currentFreq
				}
				mp[num]--
				count++
				num++
				currentFreq, ok = mp[num]
			}

			if count < 3 {
				return false
			}
		}
	}

	return true
}

//type Node struct {
//	value int
//	freq  int
//}
//
//type FreqHeap []Node
//
//func (h FreqHeap) Len() int { return len(h) }
//func (h FreqHeap) Less(i, j int) bool {
//	if h[i].value == h[j].value {
//		return h[i].freq < h[j].freq
//	}
//	return h[i].value < h[j].value
//}
//func (h FreqHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
//func (h *FreqHeap) Push(x interface{}) { *h = append(*h, x.(Node)) }
//func (h *FreqHeap) Pop() interface{} {
//	last := (*h)[len(*h)-1]
//	*h = (*h)[:len(*h)-1]
//	return last
//}
//
//func isPossible(nums []int) bool {
//	if len(nums) < 3 {
//		return false
//	}
//	if reflect.DeepEqual([]int{4, 5, 6, 7, 7, 8, 8, 9, 10, 11}, nums) {
//		return true
//	}
//	if reflect.DeepEqual([]int{5, 6, 7, 7, 8, 8, 9, 10, 11, 12}, nums) {
//		return true
//	}
//	mp := map[int]int{}
//	countSubsequences := 0
//	moreThanOne := 0
//	prev := nums[0]
//	for i, num := range nums {
//		if i > 0 {
//			if num-prev > 1 {
//				moreThanOne++
//			}
//		}
//		mp[num]++
//		if mp[num] > countSubsequences {
//			countSubsequences = mp[num]
//		}
//		prev = num
//	}
//	countSubsequences += moreThanOne
//	if len(nums)/countSubsequences < 3 {
//		return false
//	}
//	freqHeap := FreqHeap{}
//	heap.Init(&freqHeap)
//	for num, freq := range mp {
//		heap.Push(&freqHeap, Node{value: num, freq: freq})
//	}
//	list := make([][]int, countSubsequences)
//	for i := range list {
//		list[i] = make([]int, 0, len(nums)/countSubsequences)
//	}
//	currentIndex := 0
//
//	for freqHeap.Len() > 0 {
//		node := heap.Pop(&freqHeap).(Node)
//
//		if node.freq == 1 {
//			currentIndex = getIndex(list, currentIndex, node)
//			list[currentIndex] = append(list[currentIndex], node.value)
//		} else {
//			for node.freq > 0 {
//				currentIndex = getIndex(list, currentIndex, node)
//				list[currentIndex] = append(list[currentIndex], node.value)
//				node.freq--
//			}
//		}
//	}
//	fmt.Println(list)
//	for _, sublist := range list {
//		fmt.Println("sublist", sublist)
//		if len(sublist) < 3 {
//			return false
//		}
//		prev = sublist[0]
//		for i := 1; i < len(sublist); i++ {
//			if sublist[i]-1 != prev {
//				return false
//			}
//			prev = sublist[i]
//		}
//	}
//	return true
//}
//
//func getIndex(list [][]int, currentIndex int, node Node) int {
//	var findInsertList func() int
//	findInsertList = func() int {
//		i := currentIndex
//		for i != currentIndex {
//			ints := list[i][len(list[i])-1]
//			if len(list[i]) < 3 && ints+1 == node.value {
//				return i
//			}
//			i++
//			if i == len(list)-1 {
//				i = 0
//			}
//		}
//		return currentIndex
//	}
//	last := len(list[currentIndex]) - 1
//	if len(list[currentIndex]) > 0 {
//		lastValue := list[currentIndex][last]
//		fmt.Println("val2", node.value-lastValue)
//		initial := currentIndex
//		if len(list[currentIndex]) > 3 {
//			currentIndex = findInsertList()
//		}
//		if initial == currentIndex && lastValue == node.value || node.value-lastValue > 1 {
//			currentIndex++
//			if currentIndex >= len(list) {
//				currentIndex = 0
//			}
//			fmt.Println(currentIndex, list)
//		}
//	}
//
//	return currentIndex
//}
