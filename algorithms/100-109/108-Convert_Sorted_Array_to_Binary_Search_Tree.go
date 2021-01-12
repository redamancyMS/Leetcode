/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    return buildTree(nums,0,len(nums)-1)
}
func buildTree(nums []int,l int,r int) *TreeNode{
    if l > r {
        return nil
    }
    mid := (l+r)/2
    root := &TreeNode{Val:nums[mid]}
    root.Left = buildTree(nums,l,mid-1)
    root.Right = buildTree(nums,mid+1,r)
    return root
}