//因为里面数字可以重复使用，所以只能暴力搜索，但是要确定怎么查找才能不重不漏
func combinationSum(candidates []int, target int) (res [][]int) {
    
    path := []int{}
    var dfs func(c[]int,u int,target int) 
    //u表示当前访问到candidates中第几个数 
    dfs = func(c []int,u int,target int) {
        if target == 0 {
            //res = append(res,path)
            res = append(res,append([]int(nil),path...))
            return
        }
        if u == len(c) {
            return
        }
        for i := 0 ; c[u] * i <= target ; i++ {
            dfs(c,u+1,target-c[u]*i)
            path = append(path,c[u])
        }
        /*for i := 0 ; i < len(path) ; i++{
            path = append(path[:i],path[i+1:]...)
        }*/

        for i := 0 ; c[u] * i <= target ; i++ {
            path = path[:len(path)-1]
        }
    }
     dfs(candidates,0,target)
    return res
   
}