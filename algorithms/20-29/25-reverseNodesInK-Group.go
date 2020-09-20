//首先判断是否够k个元素
 //如果不够k个元素直接break
 //否则进行k个元素的交换：先内部翻转，然后考虑前后连接

func reverseKGroup(head *ListNode, k int) *ListNode {
    dummy := &ListNode{-1,nil}
    dummy.Next = head
    pre := dummy
    for true {
        q := pre
        for i := 0 ; i < k && q != nil ; i++ {
            q = q.Next
        }
        if q == nil {
            break
        }
        a := pre.Next
        b := a.Next 
        for i := 0 ; i < k-1 ; i++ {
            c := b.Next
            b.Next = a
            a = b
            b = c
        }
        c := pre.Next
        pre.Next = a
        c.Next = b
        pre = c
    }
    return dummy.Next
}