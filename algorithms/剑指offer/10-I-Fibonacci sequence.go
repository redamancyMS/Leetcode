func fib(n int) int {
    if n == 0 || n == 1 {
        return n
    }
    var dp []int = make([]int,n+1)
    dp[0] = 0
    dp[1] = 1
    for i := 2 ; i <=n ; i++ {
        dp[i] = (dp[i-1] + dp[i-2]) % (1e9 + 7)
    }
    return dp[n]

}