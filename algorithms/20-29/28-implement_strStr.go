func strStr(haystack string, needle string) int {
   m := len(haystack)
   n := len(needle)
   if n == 0 {
       return 0
   }
   if m < n {
       return -1
   }
   next := compute_next(needle)
    j := 0 
    for i := 0 ; i < m ; i++ {
        for j > 0 && haystack[i] != needle[j] {
            j = next[j-1]
        }
        if haystack[i] == needle[j] {
            j++
        }
        if j == n {
            return i - n + 1
        }
    }
    return -1

}

func compute_next(pattern string) []int {
    n := len(pattern)
    res := make([]int,n)
    k := 0
    for q := 1 ; q < n ; q++ {
        for k > 0 && pattern[q] != pattern[k] {
            k = res[k-1]
        }
        if pattern[q] == pattern[k] {
            k++
        }
        res[q] = k
    } 
    return res
}