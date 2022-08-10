package _297

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

const NilIdentifier = "0"
const NodeValueSeparator = ","

func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return NilIdentifier + NodeValueSeparator
	}
	return strconv.Itoa(root.Val) + NodeValueSeparator + this.serialize(root.Left) + this.serialize(root.Right)
}

func (this *Codec) deserialize(data string) *TreeNode {
	nodesSlice := strings.Split(data, NodeValueSeparator)
	var decodeTree func() *TreeNode
	decodeTree = func() *TreeNode {
		if nodesSlice[0] == NilIdentifier {
			nodesSlice = nodesSlice[1:]
			return nil
		}
		num, _ := strconv.Atoi(nodesSlice[0])
		nodesSlice = nodesSlice[1:]
		root := &TreeNode{num, nil, nil}
		root.Left = decodeTree()
		root.Right = decodeTree()
		return root
	}

	return decodeTree()
}
