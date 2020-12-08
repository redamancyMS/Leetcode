//这是一个简单的DP问题，dp[i]表示前i个字符可以解回去的字符串的个数
//状态计算，dp集合有两种状态，要么字符由一位数字组成，要么由两位数字组成
func numDecodings(s string) int {
    n := len(s)
    var dp []int = make([]int,n+1)
    s = " " + s
    dp[0] = 1 //当字符串为空时只有一种解码方法
    for i := 1 ; i <= n ; i++ {
        if s[i] >= '1' && s[i] <= '9'{
            dp[i] += dp[i-1]
        }
        if i > 1 {
            t := (s[i-1] - '0') * 10 + (s[i] - '0')
            if t >= 10 && t <= 26 {
                dp[i] += dp[i-2]
            }
        }
    }
    return dp[n]
}