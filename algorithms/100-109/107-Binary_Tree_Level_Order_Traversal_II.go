/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
    var ans [][]int
    var queue []*TreeNode
    if root == nil {
        return ans
    }
    queue = append(queue,root)
    for i := 0 ; len(queue) > 0 ;i++{
        q := []*TreeNode{}
        ans = append(ans,[]int{})
        for j := 0 ; j < len(queue) ; j++ {
            root = queue[j]
            ans[i] = append(ans[i],root.Val)
            if root.Left != nil {
                q = append(q,root.Left)
            }
            if root.Right != nil {
                q = append(q,root.Right)
            }
        }
        queue = q
    }
    var res [][]int
    for i := len(ans)-1 ; i >= 0 ; i--{
        res = append(res,ans[i])
    }
    return res
}