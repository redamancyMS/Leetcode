func uniquePathsWithObstacles(A [][]int) int {
     if len(A) < 1 {
         return 0
     }
     m := len(A)
     n := len(A[0])
     dp := make([][]int,m)
     for i := 0 ; i < m ; i++{
         dp[i] = make([]int,n)
     }
     for i := 0 ; i < m ; i++ {
         for j := 0 ; j < n ; j++ {
             if A[i][j] != 1 {
                 if i == 0 && j == 0 {
                     dp[i][j] = 1
                 }else{
                     if i > 0 {
                         dp[i][j] += dp[i-1][j]
                     }
                     if j > 0 {
                         dp[i][j] += dp[i][j-1]
                     }
                 }
             }
         }
     }
     return dp[m-1][n-1]
}