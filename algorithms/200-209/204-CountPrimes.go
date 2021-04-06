//朴素筛法求质数
func countPrimes(n int) int {
    if n < 2 {
        return 0
    }
    //var primes []int = make([]int,n)
    var st []bool = make([]bool,n)
    cnt := 0
    for i := 2 ; i <n ; i ++ {
        if st[i] == true{
            continue
        }
        cnt ++
        for j := i+i ; j <n ; j += i {
            st[j] = true
        }
    }
    return cnt
}