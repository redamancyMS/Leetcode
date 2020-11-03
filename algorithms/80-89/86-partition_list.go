/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//先分别用两个链表存储左右两边的数，最后将两个链表合起来
func partition(head *ListNode, x int) *ListNode {
    left := &ListNode{-1,nil}
    l := left
    right := &ListNode{-1,nil}
    r := right
    if head == nil {
        return nil
    }
    for head != nil {
        if head.Val < x {
            cur :=&ListNode{head.Val,nil}
            l.Next = cur
            l = cur
        }else{
            cur :=&ListNode{head.Val,nil}
            r.Next = cur
            r = cur
        }
        head = head.Next
    }
    l.Next = right.Next
    return left.Next
}