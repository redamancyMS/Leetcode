//对称编码，上面补零，下面补1
func grayCode(n int) []int {
    var res []int = make([]int,1)
    for n > 0 {
        for i := len(res)-1 ; i >= 0 ; i-- {
            res[i] *= 2
            res = append(res,res[i]+1)
        }
        n--
    }
    return res
}