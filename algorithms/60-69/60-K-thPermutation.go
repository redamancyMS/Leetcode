func getPermutation(n int, k int) string {
    
    var res string
    is_used := make([]bool,10)

    for i := 0 ; i < n ; i++ {
        fact := 1
        for j := 1 ; j <= n-i-1 ; j++ {
            fact *= j
        }
        for j := 1 ; j <= n ; j++ {
            if !is_used[j] {
                if fact < k {
                    k -= fact
                }else{
                    res += strconv.Itoa(j)
                    is_used[j] = true
                    break
                }  
            }
        }
    }
    return res

    
}