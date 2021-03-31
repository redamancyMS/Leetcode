func minimumTotal(triangle [][]int) int {
    if len(triangle) == 0 || len(triangle[0]) == 0 {
        return 0
    }
    var dp [][]int = make([][]int,len(triangle))
    for i := 0 ; i < len(triangle) ; i++ {
        dp[i] = make([]int,len(triangle[i]))
    }
    dp[0][0] = triangle[0][0]
    for i := 1 ; i < len(triangle) ; i++ {
        for j := 0 ; j < len(triangle[i]) ; j++ {
            if j == 0 {
                dp[i][j] = dp[i-1][j] + triangle[i][j]
            }else if j == len(triangle[i])-1{
                dp[i][j] = dp[i-1][j-1] + triangle[i][j]
            }else{
                dp[i][j] = min(dp[i-1][j-1]+triangle[i][j],dp[i-1][j]+triangle[i][j])
            }
        }
    }
    n := len(dp)
    sort.Ints(dp[n-1])
    return dp[n-1][0]
}
func min(a,b int)int {
    if a < b {
        return a
    }
    return b
}

//优化空间之后
func minimumTotal(triangle [][]int) int {
   n := len(triangle)
   if n == 0 || len(triangle[0]) == 0 {
       return 0
   }
   var dp []int = make([]int,n)
   dp[0] = triangle[0][0]
   for i := 1 ; i < n ; i ++ {
       for j := i ; j >= 0 ; j-- {
           if j == i {
               dp[j] = dp[j-1] + triangle[i][j]
           }else if j == 0{
               dp[j] = dp[j] + triangle[i][j]
           }else{
               dp[j] = min(dp[j-1],dp[j]) + triangle[i][j]
           }
       }
   }
   sort.Ints(dp)
   return dp[0]
}
func min(a,b int)int {
    if a < b {
        return a
    }
    return b
}