//用递归的方法进行判断
func isScramble(s1 string, s2 string) bool {
    if len(s1) != len(s2) {
        return false
    }
    if s1 == s2 {
        return true
    }
     s1hash, s2hash := [26]int{}, [26]int{}
  for _, ch := range s1 {
    s1hash[ch-'a']++
  }
  for _, ch := range s2 {
    s2hash[ch-'a']++
  }
  for i := 0; i < 26; i++ {
    if s1hash[i] != s2hash[i] {
      return false
    }
  }
    n := len(s1)
    for i := 1 ; i <= n-1 ; i++ {
        if isScramble(s1[0:i],s2[0:i]) && isScramble(s1[i:],s2[i:]){
            return true
        }
        if isScramble(s1[0:i],s2[n-i:]) && isScramble(s1[i:],s2[0:n-i]){
            return true
        }
    }
    return false
}