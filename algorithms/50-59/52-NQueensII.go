
//类似于全排列的问题，但是要着重判断一下每一个对角线上的问题
//向下是x轴，向右是y轴
//右斜对角线上是y-x+n(y=x+k+n)  左斜对角线上是y+x(y=-x+k)
func totalNQueens(n int) int {
   var res [][]string
   path := make([][]byte,n)
   col := make([]bool,n)//判断每一列
   zdj := make([]bool,n*2)
   udj := make([]bool,n*2)

   for i := 0 ; i < n ; i++ {
       path[i] = make([]byte,n)
       for j := 0 ; j < n ; j ++ {
           path[i][j] = '.'
       }
   }

   var dfs func(int)
   dfs = func(u int) {
       if u == n {
           ans := make([]string,n)
           for k:= 0 ; k < n ; k++ {
               ans[k] = string(path[k])
           }
           res = append(res,ans)
           return
       }
       for i := 0 ; i < n ; i++ {
           if !col[i] && !zdj[u+i] && !udj[u-i+n] {
               path[u][i] = 'Q'
               col[i] = true
               zdj[u+i] = true
               udj[u-i+n] = true
               dfs(u+1)
               path[u][i] = '.'
               col[i] = false
               zdj[u+i] = false
               udj[u-i+n] = false
           }
       }

   }
   dfs(0)
   return len(res)
    

}