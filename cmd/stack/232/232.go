package _232

type MyQueue struct {
	StackIo  []int
	StackTmp []int
	Front    int
	Flag     int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {

	if len(this.StackIo) == 0 {
		this.Front = x
	}

	this.Flag = 1
	for len(this.StackIo) != 0 {
		this.StackTmp = append(this.StackTmp, this.Pop())
	}
	this.StackTmp = append(this.StackTmp, x)

	this.Flag = 2
	for len(this.StackTmp) != 0 {
		this.StackIo = append(this.StackIo, this.Pop())
	}
	this.Flag = 0
}

func (this *MyQueue) Pop() int {

	pop := 0
	if this.Flag <= 1 && len(this.StackIo) != 0 {
		last := len(this.StackIo) - 1
		pop = this.StackIo[last]
		this.StackIo = this.StackIo[:last]
		// Update for Front(peek).
		if this.Flag == 0 && len(this.StackIo) != 0 {
			this.Front = this.StackIo[last-1]
		}

		// StackTmp mode
	} else if this.Flag == 2 && len(this.StackTmp) != 0 {
		last := len(this.StackTmp) - 1
		pop = this.StackTmp[last]
		this.StackTmp = this.StackTmp[:last]
	}

	return pop
}

func (this *MyQueue) Peek() int {
	return this.Front
}

func (this *MyQueue) Empty() bool {
	return len(this.StackIo) == 0
}
