func isValidSudoku(board [][]byte) bool {
    //按行判断
    for i := 0 ; i < 9 ; i ++ {
        is_valid := make([]bool,9)
        for j := 0 ; j <9 ; j++ {
            if board[i][j] != '.' {
                t := board[i][j] - '1'
                if is_valid[t] {
                    return false
                }
                is_valid[t] = true
            }
        }
    }
    //按列判断
    for i := 0 ; i < 9 ; i ++ {
        is_valid := make([]bool,9)
        for j := 0 ; j <9 ; j++ {
            if board[j][i] != '.' {
                t := board[j][i] - '1'
                if is_valid[t] {
                    return false
                }
                is_valid[t] = true
            }
        }
    }
    //按照小方格判断
    for i := 0 ; i < 9 ; i = i + 3 {
        for j := 0 ; j < 9 ; j = j + 3 {
            is_valid := make([]bool,9)
            for x := 0 ; x < 3 ; x++ {
                for y := 0 ; y < 3 ; y++ {
                    if board[i+x][j+y] != '.' {
                         t := board[i+x][j+y] - '1'
                        if is_valid[t] {
                            return false
                        }
                        is_valid[t] = true
                    }
                }
            }
        }
    }
    return true
}