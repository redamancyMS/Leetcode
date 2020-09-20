
func isValid(s string) bool {
    var stack []byte
    for i := 0 ; i < len(s) ; i++ {
        if s[i] == '[' || s[i] == '(' || s[i] == '{' {
            stack = append(stack,s[i])
        }else if s[i] == ')'{
            if len(stack) == 0 || stack[len(stack)-1] != '('{
                return false
            }
            stack = stack[:len(stack)-1]
        }else if s[i] == ']'{
            if len(stack) == 0 || stack[len(stack)-1] != '['{
                return false
            }
            stack = stack[:len(stack)-1]
        }else if s[i] == '}' {
            if len(stack) == 0   || stack[len(stack)-1] != '{'{
                return false
            }
            stack = stack[:len(stack)-1]
        }
    }
    if len(stack) == 0 {
        return true
    }else{
        return false 
    }
}

