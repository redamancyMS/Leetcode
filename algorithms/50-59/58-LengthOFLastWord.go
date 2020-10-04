//找最后一个单词，所以倒着找比较靠谱
func lengthOfLastWord(s string) int {
    if len(s) == 0 {
        return 0
    }
   for len(s) > 0 && s[len(s)-1] == ' '{
       s = s[:len(s)-1]
   }
    res := 0
    for i := len(s)-1 ; i >= 0 ; i-- {
        if s[i] != ' '{
            res ++
        }else{
            break
        }
    }
    return res

}