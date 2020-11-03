/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    dummy := &ListNode{-1,nil}
    dummy.Next = head
    p := dummy
    for p.Next != nil {
        q := p.Next.Next 
        for q!=nil && q.Val == p.Next.Val {
            q = q.Next
        }
        if p.Next.Next == q {
            p = p.Next
        }else{
            p.Next = q
        }
    }
    return dummy.Next

}