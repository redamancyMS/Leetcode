func plusOne(digits []int) []int {
    temp := 0
    n := len(digits)
    digits[n-1] = digits[n-1] + 1
    for i := n-1 ; i >= 0 ; i-- {
        d := digits[i] + temp
        digits[i] = d % 10
        temp = d / 10
    }
    if temp == 0 {
        return digits
    }
    var res []int 
    res = append(res,temp)
    for i := 0 ; i < n ; i++ {
        res = append(res,digits[i])
    }
    return res
}