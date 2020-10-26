//因为要去重，所以我们人为规定挑选的顺序，只能从前往后找
//dfs(n,u,start):从1到n里面选，当前还可以选u个数，当前可以从第start个数开始选
func combine(n int, k int) [][]int {
    var res [][]int
    if n < 1 || n < k {
        return res
    }
    var path []int 
    var dfs func(int,int,int) 
    dfs = func(n int,u int,start int) {
      if u == 0 {
          res = append(res,append([]int(nil),path...))
      }
      for i := start ; i <= n ; i++ {
          path = append(path,i)
          dfs(n,u-1,i+1)
          path = path[:len(path)-1]
      }
    }
    dfs(n,k,1) 
    return res
}