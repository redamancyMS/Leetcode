//这是一个图论的问题
//dp[i]表示从起点跳到下标为i的最小距离
func jump(nums []int) int {
    n := len(nums)
    var dp = make([]int,n)
    for i,j := 1,0 ; i < n ;i ++ {
        for j + nums[j] < i {
            j++
        }
        dp[i] = dp[j] + 1
    }
    return dp[n-1]
}