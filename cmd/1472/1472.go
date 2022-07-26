package _472

type Node struct {
	Val  string
	Prev *Node
	Next *Node
}

type BrowserHistory struct {
	Head    *Node
	Current *Node
	Size    int
}

func Constructor(homepage string) BrowserHistory {
	history := BrowserHistory{
		Head: &Node{
			Val:  homepage,
			Prev: nil,
			Next: nil,
		},
		Size: 0,
	}
	history.Current = history.Head
	return history
}

func (this *BrowserHistory) Visit(url string) {
	newNode := &Node{
		Val:  url,
		Prev: nil,
		Next: nil,
	}
	this.Current.Next = newNode
	newNode.Prev = this.Current
	this.Current = this.Current.Next

	this.Size++
}

func (this *BrowserHistory) Back(steps int) string {
	for this.Current.Prev != nil && steps > 0 {
		steps--
		this.Current = this.Current.Prev
	}
	return this.Current.Val
}

func (this *BrowserHistory) Forward(steps int) string {
	for this.Current.Next != nil && steps > 0 {
		steps--
		this.Current = this.Current.Next
	}
	return this.Current.Val
}

// leetcode
// google
// facebook
// leetcode
// youtube
//
