//用二进制的思想去乘，用快速幂的思想，时间复杂度是logn的
//如果是x的2²，可以化成2的10次幂，也就是2的1次幂乘以2的零次幂，以此类比
func myPow(x float64, n int) float64 {
    res := 1.0
    is_minus := n < 0
    for k := abs(n) ; k > 0 ; k >>= 1 {
        if k & 1 != 0{
            res *= x
        }
        x *= x
    }
    if is_minus {
        res = 1 / res
    }
    return res
}
func abs(n int) int{
    if n > 0 {
        return n
    }
    return -n
}