func minPathSum(grid [][]int) int {
    if len(grid) == 0 {
        return -1
    }
    m := len(grid)
    n := len(grid[0])
    var dp = make([][]int,m)
    for k := 0 ; k < m ; k++ {
        dp[k] = make([]int,n)
    }
    for i := 0 ; i < m ; i++ {
        for j := 0 ; j < n ; j++ {
            if i == 0 && j == 0 {
                dp[i][j] = grid[i][j]
            }else if i == 0 && j > 0 {
                dp[i][j] = grid[i][j] + dp[i][j-1]
            }else if j == 0 && i > 0 {
                dp[i][j] = grid[i][j] + dp[i-1][j]
            }else{
                dp[i][j] = grid[i][j] + min(dp[i][j-1],dp[i-1][j])
            }
        }
    }
    return dp[m-1][n-1]
}

func min(a,b int) int{
    if a < b {
        return a
    }
    return b
}