func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    dur := head
    nex := dur.Next
    for nex != nil {
        if nex.Val == dur.Val {
            nex = nex.Next
            dur.Next = nex
        }else{
            dur = nex
            nex = dur.Next
        }
    }
    return head
}