func solveSudoku(board [][]byte) {
    var row [9][9]bool
    var col [9][9]bool
    var cell[3][3][9]bool
    
    for i := 0 ; i < 9 ; i++ {
        for j := 0 ; j < 9 ; j++ {
            if board[i][j] != '.' {
                t := board[i][j] - '1'
                row[i][t] = true
                 col[j][t] = true
                 cell[i/3][j/3][t] = true
            }
        }
    }
    var dfs func([][]byte,int,int) bool
    dfs = func(board [][]byte,x int,y int) bool{
        if y == 9 {
            x++
            y = 0
        }
        if x == 9 {
            return true
        }
        if board[x][y] != '.' {
            return dfs(board,x,y+1)
        }
        //向九宫格里放数
        for i := byte(0) ; i < 9 ; i++ {
            if !row[x][i] && !col[y][i] && !cell[x/3][y/3][i] {
                board[x][y] = i + '1'
                row[x][i] = true 
                col[y][i] = true 
                cell[x/3][y/3][i] = true
                if dfs(board,x,y+1) {
                    return true
                }
                board[x][y] = '.'
                row[x][i] = false
                 col[y][i] = false
                  cell[x/3][y/3][i] = false
            }
        }
        return false
    }
    dfs(board,0,0)
}