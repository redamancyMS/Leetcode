func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil && l2 != nil {
        return l2
    }
    if l2 == nil && l1 != nil {
        return l1
    }
    head := &ListNode{-1,nil}
    p := head
    for l1 != nil && l2 != nil {
        if l1.Val <= l2.Val {
            q := &ListNode{l1.Val,nil}
            p.Next = q
            p = q
            l1 = l1.Next
        }else{
            q := &ListNode{l2.Val,nil}
            p.Next = q
            p = q
            l2 = l2.Next
        }
    }
    for l1 != nil {
        q := &ListNode{l1.Val,nil}
            p.Next = q
            p = q
            l1 = l1.Next
    }
    for l2 != nil {
            q := &ListNode{l2.Val,nil}
            p.Next = q
            p = q
            l2 = l2.Next
    }
    return head.Next

}

// 不重新生成链表节点
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil && l2 != nil {
        return l2
    }
    if l2 == nil && l1 != nil {
        return l1
    }
    head := &ListNode{-1,nil}
    p := head
    for l1 != nil && l2 != nil {
        if l1.Val <= l2.Val {
            p.Next = l1
            p = p.Next
            l1 = l1.Next
        }else{
            p.Next = l2
            p = p.Next
            l2 = l2.Next
        }
    }
    if l1 != nil {
            p.Next = l1  
    }
    if l2 != nil {
            p.Next = l2
    }
    return head.Next

}