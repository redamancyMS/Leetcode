/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
    var ans [][]int
    if root == nil {
        return ans
    }
    var queue []*TreeNode
    queue = append(queue,root)

    for i := 0 ; len(queue) > 0 ; i++ {
        q := []*TreeNode{}
        ans = append(ans,[]int{})
        if i == 0  {
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
        }else if i % 2 != 0{
            for j := len(queue)-1 ; j >= 0 ; j-- {
                root = queue[j]
                ans[i] = append(ans[i],root.Val)
                if root.Right != nil {
                    q = append(q,root.Right)
                }
                if root.Left != nil {
                    q = append(q,root.Left)
                }
            }
        }else{
            for j := len(queue)-1 ; j >= 0 ; j-- {
                root = queue[j]
                ans[i] = append(ans[i],root.Val)
                if root.Left != nil {
                    q = append(q,root.Left)
                }
                if root.Right != nil {
                    q = append(q,root.Right)
                }
            }
        }
        queue = q
    }
    return ans
}