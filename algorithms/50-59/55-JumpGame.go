func canJump(nums []int) bool {

    for i ,j := 0,0 ; i < len(nums) ; i++ {
       if j < i {
           return false
       }
       j = max(j,i+nums[i])
    }
    return true
   
   
}
func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}