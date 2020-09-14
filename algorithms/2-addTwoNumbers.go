给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。


//因为题目中设计的链表从头结点到最后是数字从低位到高位，所以两个数字相加的时候可以直接从头结点开始进行相加
//每个 位 上的值都是 l1.Val + l2.Val + 进位

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
     dummy := &ListNode{-1,nil}
     cur := dummy
     t := 0
     for l1 != nil || l2 != nil || t > 0 {
         if l1 != nil {
             t += l1.Val
             l1 = l1.Next
         }
         if l2 != nil {
             t += l2.Val
             l2 = l2.Next
         }
         cur.Next = &ListNode{t%10,nil}
         cur = cur.Next
         t /= 10
     }
     return dummy.Next
}
