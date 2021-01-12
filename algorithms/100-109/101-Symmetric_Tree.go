/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return leftAndRight(root.Left,root.Right)
}

func leftAndRight(l,r *TreeNode)bool{
    if l == nil && r == nil {
        return true
    }else if l == nil && r != nil {
        return false
    }else if l != nil && r == nil {
        return false
    }else if l.Val != r.Val {
        return false
    }
    return leftAndRight(l.Left,r.Right) && leftAndRight(l.Right,r.Left)
}