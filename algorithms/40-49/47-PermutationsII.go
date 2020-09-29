func permuteUnique(nums []int) (ans [][]int) {
    sort.Ints(nums)
    var path = make([]int,len(nums))
    var is_use = make([]bool,len(nums))
    var dfs func([]int,int)
    dfs = func(nums []int,u int) {
        if u == len(nums){
            ans = append(ans,append([]int(nil),path...))
            return 
        }
        for i := 0 ; i < len(nums) ; i ++ {
            if is_use[i] || i > 0 && !is_use[i-1] && nums[i] == nums[i-1] {
                continue
            }
                path[u] = nums[i]
                is_use[i] = true
                dfs(nums,u+1)
                is_use[i] = false
        }

    }
    dfs(nums,0)
    return ans
}