//用动态规划的思想来实现
//dp[i]表示所有以nums[i]结尾的区间中最大的和是多少
//dp[i]=max(dp[i-1]+nums[i],nums[i])
//时间复杂度是O(n)级别的，为了缩小空间复杂度，不再开辟新的数组空间，直接使用一个变量来代替
func maxSubArray(nums []int) int {

  var res int 
  res = math.MinInt32
  for i,last := 0,0 ; i < len(nums) ; i++ {
      last = nums[i] + max(last,0)
      res = max(res,last)
  }
  return res


}

func max(a,b int) int{
    if a > b {
        return a
    }
    return b
}