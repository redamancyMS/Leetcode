//用二进制的方法来做，假如结果k=（110010）=2¹+2^4+2^5  那么x/y可以用x-2y-2^4y-2^5y
func divide(dividend int, divisor int) int {
    var exp []int64
    is_minus := false
    if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
        is_minus = true
    }
    a := abs(int64(dividend))
    b := abs(int64(divisor))
    res := 0
    for i := b ; i <= a ; i = i + i {
        exp = append(exp,i)
    }
    for i := len(exp) - 1 ; i >= 0 ; i-- {
        if a >= exp[i] {
            res += 1 << i
            a -= exp[i]
        }
    }
    if is_minus {
        res = -res
    }
    if res > math.MaxInt32 || res < math.MinInt32 {
        res =  math.MaxInt32
    }
    return res
    
}

func abs(a int64) int64{
    if a > 0 {
        return a
    }
    return -a
}