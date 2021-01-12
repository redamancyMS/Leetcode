/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
    return inPostOrder(inorder,postorder,0,len(inorder)-1,0,len(postorder)-1)
}

func inPostOrder(inorder []int,postorder []int,l1 int,r1 int,l2 int,r2 int) (root *TreeNode){
    if l1 > r1 {
        return nil
    }
    i := l1
    for i <= r1 {
        if inorder[i] == postorder[r2]{
            root = &TreeNode{Val:inorder[i]}
            break
        }
        i++
    }
    root.Left = inPostOrder(inorder,postorder,l1,i-1,l2,r2+i-r1-1) 
    root.Right = inPostOrder(inorder,postorder,i+1,r1,r2+i-r1,r2-1)
    return root
}