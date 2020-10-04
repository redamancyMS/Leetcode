func generateMatrix(n int) [][]int {
    var res [][]int
    if n == 0 {
        return res
    }
    res = make([][]int,n)
    for i := 0 ; i < n ; i++ {
        res[i] = make([]int,n)
    }
    left := 0
    right := n-1
    up := 0
    down := n-1
    num := 1
    for true {
        for i := left ; i <= right ;i++ {
            res[up][i] = num
            num++
        }
        up++
        if up > down{
            break
        }
        for i := up ; i <= down ; i++ {
            res[i][right] = num
            num++
        }
        right--
        if right < left {
            break
        }
        for i := right ; i >=left ; i-- {
            res[down][i] = num
            num++
        }
        down--
        if down < up {
            break
        }
        for i := down ; i >= up ; i-- {
            res[i][left] = num
            num++
        }
        left++
        if left > right{
            break
        }
    }
    return res
}