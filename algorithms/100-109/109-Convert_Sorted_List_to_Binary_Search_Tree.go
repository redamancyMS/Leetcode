/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
   return buildTree(head,nil)
}
func buildTree(left,right *ListNode)  *TreeNode{
    if left == right {
        return nil
    }
    mid := findMid(left,right)
    root := &TreeNode{Val:mid.Val}
    root.Left = buildTree(left,mid)
    root.Right = buildTree(mid.Next,right)
    return root

}
func findMid(l,r *ListNode) *ListNode{
    fast,slow := l,l
    for fast != r && fast.Next != r {
        fast = fast.Next.Next
        slow = slow.Next
    }
    return slow
}