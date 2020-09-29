func combinationSum2(candidates []int, target int) (ans [][]int) {
    sort.Ints(candidates)
    var path []int
   
    var dfs func([]int,int,int)
    dfs = func(c []int,u int,target int) {
        if target == 0 {
            ans = append(ans,append([]int(nil),path...))
            return 
        }
        if u == len(c) {
            return 
        }
        k := u+1
        for k < len(c) && c[k]==c[u]{
            k++
        }
        cnt := k-u
        for i := 0 ; c[u]*i <= target && i <= cnt ; i++ {
            dfs(c,k,target-c[u]*i)
            path = append(path,c[u])
        }
         for i := 0 ; c[u]*i <= target && i <= cnt ; i++ {
            path = path[:len(path)-1]
        }
    }
    dfs(candidates,0,target)
    return ans
}