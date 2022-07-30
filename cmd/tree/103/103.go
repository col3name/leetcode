*
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNodee
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
    var result [][]int
    
    stack := make ([]*TreeNode,0)
    stack = append(stack, root)
    vector := 1
    for len(stack) > 0 {
        var tmp []*TreeNode
        var level []int
        for i := len(stack) - 1; i >= 0; i-- {       
            node := stack[i]
            if node == nil {
                continue
            }
            if vector == 1 {
                tmp = append(tmp, node.Left, node.Right)
            } else {
                tmp = append(tmp, node.Right, node.Left)
            }
                 
            level = append(level, node.Val)
                                              
            
                                      
        }
        vector = -vector
        stack = tmp
        if len(level) > 0 {
            result = append(result, level)   
        }
    }
    return result
}
