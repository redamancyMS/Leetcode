func solveSudoku(board [][]byte) {
   var row [9][9]bool //每一行判断1-9是否使用
   var col [9][9]bool //每一列判断哪1-9是否使用
   var box [3][3][9]bool  //每个小方格判断1-9是否使用

    for i := 0 ; i < 9 ; i++ {
        for j := 0 ; j < 9 ; j++ {
            if board[i][j] != '.' {
                t := board[i][j] - '1'
                row[i][t] = true
                col[j][t] = true
                box[i/3][j/3][t] = true
            }
        }
    }
    var dfs func([][]byte,int,int) bool
    dfs = func(nums [][]byte,x int,y int) bool {
        if y==9 {
            y = 0
            x++
        }
        if x == 9 {
            return true
        }
        if nums[x][y] !='.'{
           return dfs(nums,x,y+1)
        }
        //这个是用来判断当前放哪个数字的
        for i := byte(0) ; i < 9 ; i++{
            if !row[x][i] && !col[y][i] && !box[x/3][y/3][i] {
                nums[x][y] = i + '1'
                row[x][i] = true
                col[y][i] = true
                box[x/3][y/3][i] = true
                if dfs(nums,x,y+1) { // 找到了答案
                    return true
                }
                //否则就恢复现场
                nums[x][y] = '.'
                row[x][i] = false
                col[y][i] = false
                box[x/3][y/3][i] = false
            }
        }
        //没有找到答案
        return false
    }
   dfs(board,0,0)
}
