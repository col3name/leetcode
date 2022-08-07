package _3

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	var resultHead *ListNode
	var cur *ListNode
	cur = &ListNode{}
	resultHead = cur
	cur.Next = &ListNode{}
	localList := lists
	for cur != nil {
		cur.Next, localList = findMin(localList)
		cur = cur.Next
	}

	return resultHead.Next
}

func findMin(lists []*ListNode) (*ListNode, []*ListNode) {
	var min *ListNode
	var minValue int
	minValue = 10000000000
	minIndex := -1
	var localList []*ListNode
	localList = lists
	for i := 0; i < len(localList); i++ {
		node := localList[i]
		if node != nil {
			if node.Val < minValue {
				minIndex = i
				minValue = node.Val
			}
		}
	}
	if minIndex == -1 {
		return nil, localList
	}
	min = &ListNode{
		Val:  minValue,
		Next: nil,
	}
	node := localList[minIndex]

	localList[minIndex] = node.Next
	return min, localList
}
