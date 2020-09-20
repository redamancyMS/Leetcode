func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }
    if len(lists) == 1 {
        return lists[0]
    } 
    p := lists[0]
    for i := 1 ; i < len(lists) ; i++ {
        p = merge2List(p,lists[i])
    }
    return p
}

func merge2List(l1 *ListNode,l2 *ListNode) *ListNode {
    dummy := &ListNode{-1,nil}
    pre := dummy

    for l1 != nil && l2 != nil {
        if l1.Val <= l2.Val {
            pre.Next = l1 
            pre = l1
            l1 = l1.Next
        }else{
            pre.Next = l2
            pre = l2
            l2 = l2.Next
        }
    }
    if l1 != nil {
        pre.Next = l1
    }
    if l2 != nil {
        pre.Next = l2
    }
    return dummy.Next

}