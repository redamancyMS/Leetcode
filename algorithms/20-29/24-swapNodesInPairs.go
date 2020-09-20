func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    dummy := &ListNode{-1,nil}
    dummy.Next = head
    pre := dummy
    for pre.Next != nil && pre.Next.Next != nil{
       node1 := pre.Next
       node2 := node1.Next
       next := node2.Next
       pre.Next = node2
       node1.Next = next
       node2.Next = node1
       pre = node1
    }
    return dummy.Next

}