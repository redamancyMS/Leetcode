//这是一个动态规划的问题
//dp[i][j] 表示从word1[1...i]变到word2[1...j]所需要的最少操作为多少
func minDistance(word1 string, word2 string) int {
   n := len(word1)
   m := len(word2)
   word1 = " " + word1
   word2 = " " + word2
   var dp [][]int = make([][]int,n+1)
   for i := 0 ; i <= n ; i++ {
       dp[i] = make([]int,m+1)
   }
   for i := 0 ; i <= n ; i++ {
       dp[i][0] = i
   }
   for i := 1 ; i <= m ; i++ {
       dp[0][i] = i
   }

   for i := 1 ; i <= n ; i++ {
       for j := 1 ; j <= m ; j++ {
           //增加或者删除操作时,dp[i-1][j]是删除操作，dp[i][j-1]是增加操作
           dp[i][j] = min(dp[i-1][j]+1,dp[i][j-1]+1)
           //修改操作时需要判断两值是否相等
           t := 0
           if word1[i] != word2[j] {
               t = 1
           }
           dp[i][j] = min(dp[i][j],dp[i-1][j-1]+t)
       }
   }
   return dp[n][m]
}

func min(a,b int) int {
    if a < b {
        return a
    }
    return b

}