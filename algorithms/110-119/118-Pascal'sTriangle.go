func generate(numRows int) [][]int {
    var res [][]int
    if numRows == 0 {
        return res
    }
    if numRows == 1 {
        row := []int{1}
        res = append(res,row)
    }else{
        res = append(res,[]int{1})
        for i := 2 ; i <= numRows; i++ {
            row := make([]int,i)
            row[0] = 1
            for j := 1 ; j < i-1 ; j++{
                row[j]=res[i-2][j-1]+res[i-2][j]
            }
            row[i-1] = 1
            res = append(res,row)
        }
    }
    return res
}