func exist(board [][]byte, word string) bool {
    
    var dx []int = []int{-1,0,1,0} 
    var dy []int = []int{0,1,0,-1} 

    var dfs func([][]byte,string,int,int,int) bool
    dfs = func(board [][]byte,word string,u int,x int,y int) bool{
        if board[x][y] != word[u] {
            return false
        }
        if u == len(word)-1 {
            return true
        }
        t := board[x][y]
        board[x][y] = '.'
        for i := 0 ; i < 4 ; i++ {
            a := x + dx[i]
            b := y + dy[i]
            if a < 0 || a >= len(board) ||  b < 0 || b >= len(board[0]) || board[a][b] == '.' {continue}
            if dfs(board,word,u+1,a,b){return true}   
        }
        board[x][y] = t 
        return false
    }
    for i := 0 ; i < len(board) ; i++ {
        for j := 0 ; j < len(board[0]) ; j++ {
            if dfs(board,word,0,i,j) {
                return true
            }
        }
    }
    return false
}