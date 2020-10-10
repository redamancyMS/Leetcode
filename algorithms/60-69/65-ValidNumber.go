//1、去掉前后空格
//2、e的前后如果是空则为false
//3、e的后面不能有‘.’,e和.最多出现一次
//4、+-不能连续出现多于1个，+-的后面必须有数字 
func isNumber(s string) bool {
    if len(s) == 0 {
        return false
    }
    l := 0
    r := len(s)-1
    //先把前后空格去掉
    for l <= r && s[l] == ' ' {
        l++
    }
    for l <= r && s[r] == ' '{
        r--
    }
    if l > r {
        return false
    }
    s = s[l:r+1]
    //去掉正负号
    if s[0] == '+' || s[0] == '-' {
        s = s[1:]
    }
    if len(s) == 0 {
        return false
    }
    //如果只有.或者e前后有.
    if s[0] == '.' && (len(s) == 1 || s[1] == 'e' || s[1] == 'E'){
        return false
    }
    dot,e := 0,0
    //开始判断剩下的字符
    for i := 0 ; i < len(s) ; i++{
        if s[i] == '.' {
            if dot > 0 || e > 0 {
                return false
            }
            dot++
        }else if s[i] == 'e' || s[i] == 'E' {
            if i==0 || i+1 == len(s) || e > 0 {
                return false
            }
            if s[i+1] == '+' || s[i+1] == '-' {
                if i+2 == len(s) {
                    return false
                }
                i++
            }
            e++ 
        }else if(s[i] < '0' || s[i] > '9'){
            return false
        }
    }
    return true
}