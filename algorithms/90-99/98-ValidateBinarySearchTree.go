/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
 
    var stack []*TreeNode
    var inorder []int
    for root != nil || len(stack)>0 {
        if root != nil {
            stack = append(stack,root)
            root = root.Left
        }else{
            root = stack[len(stack)-1]
            inorder = append(inorder,root.Val)
            stack = stack[:len(stack)-1]
            root = root.Right
        }
    }
    for i := 1 ; i < len(inorder) ; i++ {
        if inorder[i-1] >= inorder[i]{
            return false
        }
    }
    return true
}