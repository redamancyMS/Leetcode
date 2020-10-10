//方法一
var dp = make([][101]int, 101)//dp[m][n] 对于一个m*n网格的有几种路径

func uniquePaths(m int, n int) int {
    if m <= 0 && n <= 0 {
        return 0
    }else if m == 1 || n == 1 {
        return 1
    }else if m == 2 && n == 2 {
        return 2
    }else if m == 2 && n == 3 || m == 3 && n == 2 {
        return 3
    }
    if dp[m][n] > 0 {
        return dp[m][n]
    }
    dp[m-1][n] = uniquePaths(m-1,n)
    dp[m][n-1] = uniquePaths(m,n-1)
    dp[m][n] = dp[m-1][n] + dp[m][n-1]
    return dp[m][n]
}

//方法二
func uniquePaths(m int, n int) int {
     dp := make([][]int,m)
     for i := 0 ; i < m ; i++{
         dp[i] = make([]int,n)
     }
     for i := 0 ; i < m ; i++ {
         for j := 0 ; j < n ; j++ {
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
     return dp[m-1][n-1]
}