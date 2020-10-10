/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil {
        return head
    }
    var tail *ListNode
    n := 0
    for p := head ; p != nil ; p = p.Next {
        n++
        tail = p
    }
    k %= n
    ptr := head
    for i := 0 ; i < n-k-1 ; i++ {
        ptr = ptr.Next
    }
    tail.Next = head
    head = ptr.Next
    ptr.Next = nil
    return head
}