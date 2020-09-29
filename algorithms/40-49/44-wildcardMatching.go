//dp[i][j]表示s[1..i]和p[1..j]是否匹配
func isMatch(s string, p string) bool {
    m := len(s)
    n := len(p)
    s = " " + s
    p = " " + p
     dp := make([][]bool,m+1)
     for i := 0 ; i < m+1 ; i++ {
         dp[i] = make([]bool,n+1)
     }
     dp[0][0] = true

     for i := 0 ; i < m+1 ; i++ {
         for j := 1 ; j < n+1 ; j++ {
             if p[j] != '*' {
                 dp[i][j] = i > 0 && dp[i-1][j-1] && (s[i]==p[j] || p[j]=='?')
             }else{
                 dp[i][j] = dp[i][j-1] || i > 0 && dp[i-1][j]
             }
         }
     }
     return dp[m][n]
}