// 典型的滑动窗口思想
func minWindow(s string, t string) string {
   need := make(map[byte]int)
   window := make(map[byte]int)

   for i := 0 ; i < len(t) ; i++ {
       need[t[i]]++
   }
   l,r := 0,0
   valid := 0
   start := 0
   length := math.MaxInt32
   for r < len(s) {
       c := s[r]
       r++
       if _,ok := need[c];ok {
           window[c]++
           if window[c] == need[c]{
               valid++
           }
       }
       for valid == len(need) {
           if r-l < length {
               start = l
               length = r-l
           }
           d := s[l]
           l++ 
           if _,ok := need[d];ok {
               if window[d]==need[d]{
                   valid--
               }
               window[d]--
           }
       }
   }
   if length != math.MaxInt32 {
       return s[start:start+length]
   }
   return ""

}