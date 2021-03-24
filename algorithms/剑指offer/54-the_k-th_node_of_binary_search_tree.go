//遍历之后可以直接取出
func kthLargest(root *TreeNode, k int) int {
    var stack []*TreeNode
    var ans []int
    for root!= nil || len(stack) > 0 {
        if root != nil {
            stack = append(stack,root)
            root = root.Left
        }else{
            root = stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            ans = append(ans,root.Val)
            root = root.Right
        }
    }
    return ans[len(ans)-k]
}