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

func (c *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "nil,"
	}
	return strconv.Itoa(root.Val) + "," + c.serialize(root.Left) + c.serialize(root.Right)
}

func (c *Codec) deserialize(data string) *TreeNode {
	array := strings.Split(data, ",")
	var decodeTree func() *TreeNode
	decodeTree = func() *TreeNode {
		if array[0] == "nil" {
			array = array[1:]
			return nil
		}
		number, _ := strconv.Atoi(array[0])
		array = array[1:]
		root := &TreeNode{number, nil, nil}
		root.Left = decodeTree()
		root.Right = decodeTree()
		return root
	}

	return decodeTree()
}
