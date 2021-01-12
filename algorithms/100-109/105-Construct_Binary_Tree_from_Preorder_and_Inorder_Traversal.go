/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    return preInorder(preorder,inorder,0,len(preorder)-1,0,len(inorder)-1)
}

func preInorder(preorder []int,inorder []int,l1 int,r1 int,l2 int,r2 int) (root *TreeNode){
    if l1 > r1 {
        return nil
    }
    i := l2
    for i <= r2 {
        if preorder[l1] == inorder[i]{
            break
        }
        i++
    }
    root = &TreeNode{Val:inorder[i]}
    root.Left = preInorder(preorder,inorder,l1+1,l1+i-l2,l2,i-1) 
    root.Right = preInorder(preorder,inorder,l1+i-l2+1,r1,i+1,r2)
    return root
}