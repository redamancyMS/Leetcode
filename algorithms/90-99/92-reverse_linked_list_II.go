/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
    dummy := &ListNode{-1,nil}
    dummy.Next = head
    a := dummy
    for i := 0 ; i < m - 1 ; i++ {
        a = a.Next
    }
    b := a.Next
    c := b.Next
    for i := 0 ; i < n-m ;i ++ {
        t := c.Next
        c.Next = b
        b = c
        c = t
    }
    a.Next.Next = c
    a.Next = b
    return dummy.Next
}