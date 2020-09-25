func countAndSay(n int) string {
    s := "1"
    for i := 0 ; i < n-1 ; i++ {
        var t string
        j :=0
        for j < len(s) {
            k := j+1
            for k < len(s) && s[k]==s[j] {
                k++
            }
            t = t + strconv.Itoa(k-j) + string(s[j])
            j = k
        }
        s = t
    }
    return s
}