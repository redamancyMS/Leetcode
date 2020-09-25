func findSubstring(s string, words []string) []int {
    var res []int
    if len(words) == 0 {
        return res
    }
    n := len(s)
    m := len(words)
    w := len(words[0])
    need := make(map[string]int)
    for _,word := range words {
        need[word] ++
    }
    for i:= 0 ; i < w ; i++ {
        window := make(map[string]int)
        cnt := 0
        for j:=i ; j + w <=n ; j += w {
            if j >= i + m*w {
                word := s[j-m*w:j-m*w+w]
                window[word]--
                if window[word] < need[word] {
                    cnt --
                }
            }
            word := s[j:j+w]
            window[word]++
            if window[word] <= need[word] {
                cnt ++
            }
            if cnt == m {
                res = append(res,j-(m-1)*w)
            }
        }
    }
    return res
}