//动态规划，可以类比最长公共子序列的题目
//dp[i][j]表示所有由s1[1~i]和s2[1~j]交错形成s3[1~i+j]的方案是否为0
func isInterleave(s1 string, s2 string, s3 string) bool {
    n := len(s1)
    m := len(s2)
    if len(s3) != m+n {
        return false
    }
    s1 = " "+s1
    s2 = " "+s2
    s3 = " "+s3
    var dp [][]bool = make([][]bool,n+1)
    for i := 0 ; i <= n ; i++ {
        dp[i] = make([]bool,m+1)
    }
    for i := 0 ; i <= n ; i ++ {
        for j := 0 ; j <= m ; j++{
            if i == 0 && j == 0 {
                dp[i][j] = true
            }
            if i > 0 && s1[i] == s3[i+j]{
                dp[i][j] = dp[i-1][j]
            }
            if j > 0 && s2[j] == s3[i+j]{
                dp[i][j] = dp[i][j] || dp[i][j-1]
            }
        }
    }
    return dp[n][m]

}