/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 //类似于层次遍历
func minDepth(root *TreeNode) int {
   
    if root == nil {
        return 0
    }
    var queue []*TreeNode
    var depth int = 1
    queue = append(queue,root)
    for len(queue) > 0 {
        num := len(queue)
        for num > 0 {
            node := queue[0]
            queue = queue[1:]
            if node.Left == nil && node.Right == nil {
                return depth
            }
            if node.Left != nil {
                queue = append(queue,node.Left)
            }
            if node.Right != nil{
                queue = append(queue,node.Right)
            }
            num--
        }
        depth ++
    }
    return depth
}