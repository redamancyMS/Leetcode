//1、需要记录每个位置填什么  []int
//2、 需要记录当前的位置  u
//3、记录当前哪些数已经用过  []bool
func permute(nums []int) [][]int {
    var res [][]int
    n := len(nums)
    var path []int = make([]int,n)
    var is_use []bool = make([]bool,n)

    var dfs func([]int,int)
    dfs = func(nums []int,u int) {
        if u == len(nums) {
           res = append(res,append([]int(nil),path...))
            return
        }
        for i := 0 ; i < len(nums) ; i++ {
            if is_use[i] == false {
                path[u] = nums[i]
                is_use[i] = true
                dfs(nums,u+1)
                is_use[i] = false
            }  
        }
    }
    dfs(nums,0)
    return res

}