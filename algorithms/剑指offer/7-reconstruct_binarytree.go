/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    root := buildPreInoder(preorder,inorder,0,len(preorder)-1,0,len(inorder)-1)
    return root
}
func buildPreInoder(preorder []int,inorder []int,l1 int,r1 int,l2 int,r2 int) (root *TreeNode){
    if l1 > r1 {
        return nil
    }
    var i int
    for i = l2 ; i <= r2 ; i ++ {
        if inorder[i] == preorder[l1] {
            break
        }
    }
    root = &TreeNode{Val:inorder[i]}
    root.Left = buildPreInoder(preorder,inorder,l1+1,l1+i-l2,l2,i-1)
    root.Right = buildPreInoder(preorder,inorder,l1+i-l2+1,r1,i+1,r2)
    return root
}