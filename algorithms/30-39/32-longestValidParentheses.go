func longestValidParentheses(s string) int {
    if len(s) == 0 {
        return 0
    }
    var stack []int
    res := 0
    for i,start:=0,-1 ; i < len(s) ; i++ {
        if s[i] == '(' {
            //将括号的下标放入
            stack = append(stack,i)
        }else{
            if len(stack) != 0 {
                stack = stack[:len(stack)-1]
                if len(stack) != 0 {
                    res = max(res,i-stack[len(stack)-1])
                }else{
                    res = max(res,i-start)
                }
            }else{
                start = i
            }
        }
    }
    return res
        
}
func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}