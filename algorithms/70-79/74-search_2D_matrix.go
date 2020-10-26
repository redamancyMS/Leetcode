func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix)==0 || len(matrix[0])==0 {
        return false
    }
    m := len(matrix)
    n := len(matrix[0])

    for i := 0 ; i < m ; i++ {
        if target > matrix[i][n-1]{continue}
        l,r := 0,n-1
        for l < r {
            mid := (l+r)/2
            if matrix[i][mid] == target {
                return true
            }
            if target > matrix[i][mid]{
                l = mid +1
            }else{
                r = mid
            }
        }
        if matrix[i][r] == target{
            return true
        }      
    }
    return false
}