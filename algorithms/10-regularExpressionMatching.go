正则表达式匹配

给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。

'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。

说明:

s 可能为空，且只包含从 a-z 的小写字母。
p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。
示例 1:

输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。
示例 2:

输入:
s = "aa"
p = "a*"
输出: true
解释: 因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
示例 3:

输入:
s = "ab"
p = ".*"
输出: true
解释: ".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
示例 4:

输入:
s = "aab"
p = "c*a*b"
输出: true
解释: 因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。
示例 5:

输入:
s = "mississippi"
p = "mis*is*p*."
输出: false


//很明显要使用动态规划来解决此问题，重要的是要找到状态和状态转移公式

func isMatch(s string, p string) bool {
     m := len(s)
     n := len(p)
     s = " " + s
     p = " " + p

     dp := make([][]bool,m+1)
     for i := 0 ; i < m + 1 ; i++ {
         dp[i] = make([]bool,n+1)
     }
     dp[0][0] = true

     for i := 0 ; i <= m ; i++ {
         for j := 1 ; j <= n ; j++ {
             //如果当前字符的下一个字符是*，那么当前字符无法作为单独的个体判断，必须和*作为一个整体进行判断
             if j+1 <= n && p[j+1] == '*'{
                 continue
             }
             if i > 0 && p[j] != '*' {
                 dp[i][j] = dp[i-1][j-1] && (s[i]==p[j] || p[j]=='.')
             }else if p[j] == '*' {
                 dp[i][j] = dp[i][j-2] || i > 0 && dp[i-1][j] && (s[i]==p[j-1] || p[j-1]=='.')
             }
                 
             
         }
     }
     return dp[m][n]
}
