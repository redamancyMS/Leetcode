//原地算法，可以用第一行来记录每一列是否包含零元素，用第一列来记录每一行是否包含零元素，同时再用两个额外的变量来记录第一行和第一列是否包含零元素
func setZeroes(matrix [][]int)  {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return
    }
    m,n := len(matrix),len(matrix[0])
    r,c := 1,1 //1表示没有零，0则表示有
    for i := 0 ; i < n ; i++ {
        if matrix[0][i] == 0 {r = 0}
    }
    for i := 0 ; i < m ; i ++ {
        if matrix[i][0] == 0 {c = 0}
    }

     //判断每一列
    for i := 1 ; i < m ; i ++ {
        for j := 0 ; j < n ; j++ {
            if matrix[i][j] == 0 {matrix[i][0]=0}
        }
    }

    //判断每一行
    for i := 1 ; i < n ; i ++ {
        for j := 0 ; j < m ; j++ {
            if matrix[j][i] == 0 {matrix[0][i]=0}
        }
    }
    for i := 1 ; i < n ; i ++ {
        if matrix[0][i] == 0 {
            for j := 0 ; j < m ; j++ {
                matrix[j][i] = 0
            }
        }
    }
    for i := 1 ; i < m ; i ++ {
        if matrix[i][0] == 0 {
            for j := 0 ; j < n ; j++ {
                matrix[i][j] = 0
            }
        }
    }
    if r==0 {
        for i := 0 ; i < n ; i++{
            matrix[0][i] = 0
        }
    }
    if c == 0 {
        for i := 0 ; i < m ; i++ {
            matrix[i][0] = 0
        }
    }
   

}