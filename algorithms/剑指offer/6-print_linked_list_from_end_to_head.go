/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
    var nums []int
    for head != nil {
        nums = append(nums,head.Val)
        head = head.Next
    }
    for i,j := 0,len(nums)-1;i < j ; i,j= i+1,j-1 {
        swap(&nums[i],&nums[j])
    }
    return nums
}
func swap(a,b *int) {
    *a,*b = *b,*a
}