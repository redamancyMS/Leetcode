func removeNthFromEnd(head *ListNode, n int) *ListNode {
    
    pre := &ListNode{-1,nil}
    cur := &ListNode{0,nil}
    pre.Next = head
    cur = pre
    fast := pre.Next
    slow := pre.Next
    for n > 0 {
        fast = fast.Next
        n--
    }

    for fast != nil {
        fast = fast.Next
        slow = slow.Next
        cur = cur.Next
    }
    cur.Next = slow.Next
    return pre.Next
}