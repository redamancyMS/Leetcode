func replaceSpace(s string) string {
    if len(s) == 0 {
        return s
    }
    r := []rune(s)
    var res string = ""
    for i := 0 ; i < len(r) ; i++ {
        if r[i] == ' '{
            res += "%20"
        }else{
            res += string(r[i])
        }
    }
    return res  
}